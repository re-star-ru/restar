package main

import (
	"context"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-pkgz/rest"
	"github.com/go-pkgz/rest/logger"
	"github.com/google/uuid"
	"github.com/grandcat/zeroconf"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"restar/configs"
	"restar/pkg/diagnostic"
	"restar/pkg/interceptors"
	"restar/pkg/user"
)

func main() {
	config := configs.NewConfig()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	s := zerolog.NewConsoleWriter()
	s.Out = os.Stderr
	s.TimeFormat = time.StampMilli

	log.Logger = log.Output(s)

	go runHealth(config)
	run(config)
}

func run(c configs.Config) {
	// setup logging and recovery
	srv := grpc.NewServer(grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
		interceptors.ZerologUnaryServerInterceptor(),
		grpcrecovery.UnaryServerInterceptor(),
	)))
	reflection.Register(srv)

	listen, err := net.Listen("tcp", c.Host)
	if err != nil {
		log.Fatal().Err(err).Msg("cant listen grpc service")
	}

	user.RegisterService(srv, user.NewUserUsecase())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	conn, err := pgx.Connect(ctx, c.Postgres)
	if err != nil {
		log.Error().Err(err).Msg("cant connect to postgres")
		return
	}

	drepo := diagnostic.NewPostgresRepo(conn)
	diagnostic.RegisterService(srv, diagnostic.NewUsecase(drepo))

	log.Info().Msgf("restar service listen at %s", c.Host)
	log.Error().Err(srv.Serve(listen)).Msg("cant serve grpc service")
}

/////////////////////////// service discovery and else

func runHealth(c configs.Config) {
	instanceID := uuid.NewString()
	setupPeers(instanceID)

	r := chi.NewRouter()
	r.Use(rest.Ping, logger.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { rest.RenderJSON(w, "Instance id: "+instanceID) })

	log.Print("health listen at:", c.DiscoveryHost)

	timeout := time.Second * 5
	server := &http.Server{
		Addr:              c.DiscoveryHost,
		Handler:           r,
		ReadTimeout:       timeout,
		ReadHeaderTimeout: timeout,
		WriteTimeout:      timeout,
		IdleTimeout:       timeout,
	}

	log.Fatal().Err(server.ListenAndServe()).Send()
}

func setupPeers(id string) {
	log.Print("id:", id)

	port := openRandomPort()
	go registerPeer(id, port)
	go listenForNewPeers(id)
}

func registerPeer(id string, port int) {
	server, err := zeroconf.Register(
		id,
		"_workstation._tcp",
		"local.",
		port,
		[]string{"txtv=0", "lo=1", "la=2"},
		nil,
	)
	if err != nil {
		panic(err)
	}

	defer server.Shutdown()

	select {}
}

func listenForNewPeers(self string) {
	resolver, err := zeroconf.NewResolver()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize resolver")
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			if entry.Instance == self {
				continue
			}

			log.Printf("Found peer: id: %v, port: %v, ips: %v", entry.Instance, entry.Port, entry.AddrIPv4)
		}
	}(entries)

	err = resolver.Browse(context.Background(), "_workstation._tcp", "local.", entries)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to browse")
	}

	select {}
}

func openRandomPort() int {
	rand.Seed(time.Now().UnixNano())

	start := 20000
	end := 22000

	return rand.Intn(end-start+1) + start //nolint:gosec
}
