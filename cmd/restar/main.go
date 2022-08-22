package main

import (
	"context"
	"math/rand"
	"net/http"
	"os"
	"restar/configs"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-pkgz/rest"
	"github.com/go-pkgz/rest/logger"
	"github.com/google/uuid"
	"github.com/grandcat/zeroconf"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config := configs.NewConfig()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMilli})

	instanceID := uuid.NewString()
	setupPeers(instanceID)

	r := chi.NewRouter()
	r.Use(rest.Ping, logger.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { rest.RenderJSON(w, "Instance id: "+instanceID) })

	log.Print("listen at:", config.Host)
	log.Fatal().Err(http.ListenAndServe(config.Host, r)).Send()
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
