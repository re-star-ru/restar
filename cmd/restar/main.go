package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/grandcat/zeroconf"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	log.Println("Hello Restar")

	id := uuid.NewString()

	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World, id: " + id))
	})

	log.Println("id:", id)

	port := openRandomPort()
	host := fmt.Sprintf(":%d", port)

	go registerPeer(id, port)
	go listenForNewPeers(id)

	log.Println("listen at:", host)
	log.Println(http.ListenAndServe(host, m))

	select {}
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
		log.Fatalf("Failed to initialize resolver: %s", err)
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
		log.Fatalln("Failed to browse:", err.Error())
	}

	select {}
}

func openRandomPort() int {
	rand.Seed(time.Now().UnixNano())

	start := 20000
	end := 22000

	return rand.Intn(end-start+1) + start //nolint:gosec
}
