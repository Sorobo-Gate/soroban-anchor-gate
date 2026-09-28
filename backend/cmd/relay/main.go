package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ready","service":"soroban-anchor-gate-relay"}`))
}

func main() {
	http.HandleFunc("/health", healthHandler)
	port := ":8080"
	fmt.Printf("SorobanAnchor Gate Relay listening on %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server startup failed: %v", err)
	}
}
