package handler

import (
	"context"
	"net/http"

	"petstore/internal/api"
)

func (s *PetStoreHandler) FindPetByID(ctx context.Context, params api.FindPetByIDParams) (*api.Pet, error) {
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
