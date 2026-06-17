package main

import (
	"log"
	"net/http"

	"petstore/internal/api"
	"petstore/internal/server"
)

func main() {
	handler := server.NewPetStoreServer()
	srv, err := api.NewServer(handler)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
