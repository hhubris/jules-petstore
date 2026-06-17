package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"petstore/internal/server"
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := server.Run(ctx); err != nil {
		log.Printf("Server failed: %v", err)
		os.Exit(1)
	}
}
