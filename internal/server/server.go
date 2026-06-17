package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"petstore/internal/api"
	"petstore/internal/handler"
)

// Run creates the server and starts listening, observing the given context for shutdown
func Run(ctx context.Context) error {
	h := handler.New()
	srv, err := api.NewServer(h)
	if err != nil {
		return err
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: srv,
	}

	go func() {
		log.Println("Server listening on :8080")
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen failed: %v", err)
		}
	}()

	// Wait for context cancellation
	<-ctx.Done()
	log.Println("Shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return httpServer.Shutdown(shutdownCtx)
}
