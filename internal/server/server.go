package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"petstore/internal/api"
)

// Run creates the server and starts listening, observing the given context for shutdown
func Run(ctx context.Context) error {
	handler := NewPetStoreServer()
	srv, err := api.NewServer(handler)
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

var _ api.Handler = (*PetStoreServer)(nil)

type PetStoreServer struct {
	mu     sync.Mutex
	pets   map[int64]*api.Pet
	nextID int64
}

func NewPetStoreServer() *PetStoreServer {
	return &PetStoreServer{
		pets:   make(map[int64]*api.Pet),
		nextID: 1,
	}
}

func (s *PetStoreServer) AddPet(ctx context.Context, req *api.NewPet) (*api.Pet, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pet := &api.Pet{
		ID:   s.nextID,
		Name: req.Name,
		Tag:  req.Tag,
	}
	s.pets[s.nextID] = pet
	s.nextID++

	return pet, nil
}

func (s *PetStoreServer) DeletePet(ctx context.Context, params api.DeletePetParams) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.pets[params.ID]; !ok {
		return &api.ErrorStatusCode{
			StatusCode: http.StatusNotFound,
			Response: api.Error{
				Code:    int32(http.StatusNotFound),
				Message: "Pet not found",
			},
		}
	}

	delete(s.pets, params.ID)
	return nil
}

func (s *PetStoreServer) FindPetByID(ctx context.Context, params api.FindPetByIDParams) (*api.Pet, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pet, ok := s.pets[params.ID]
	if !ok {
		return nil, &api.ErrorStatusCode{
			StatusCode: http.StatusNotFound,
			Response: api.Error{
				Code:    int32(http.StatusNotFound),
				Message: "Pet not found",
			},
		}
	}

	return pet, nil
}

func (s *PetStoreServer) FindPets(ctx context.Context, params api.FindPetsParams) ([]api.Pet, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []api.Pet
	for _, pet := range s.pets {
		if params.Tags != nil {
			// Basic filtering if tag is provided
			match := false
			for _, tag := range params.Tags {
				if pet.Tag.IsSet() && pet.Tag.Value == tag {
					match = true
					break
				}
			}
			if !match {
				continue
			}
		}
		result = append(result, *pet)
	}

	// Sort by ID to make it deterministic
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	if params.Limit.IsSet() {
		limit := int(params.Limit.Value)
		if len(result) > limit {
			result = result[:limit]
		}
	}

	return result, nil
}

func (s *PetStoreServer) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	if apiErr, ok := err.(*api.ErrorStatusCode); ok {
		return apiErr
	}

	return &api.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		},
	}
}
