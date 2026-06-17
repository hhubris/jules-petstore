package handler

import (
	"context"
	"net/http"

	"petstore/internal/api"
)

func (s *PetStoreHandler) DeletePet(ctx context.Context, params api.DeletePetParams) error {
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
