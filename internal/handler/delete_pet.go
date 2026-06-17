package handler

import (
	"context"
	"errors"
	"net/http"

	"petstore/internal/api"
	"petstore/internal/domain"
	"petstore/internal/storage"
)

func (s *PetStoreHandler) DeletePet(ctx context.Context, params api.DeletePetParams) error {
	input := domain.DeletePetInput{
		ID: params.ID,
	}

	_, err := s.store.DeletePet(ctx, input)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return &api.ErrorStatusCode{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Code:    int32(http.StatusNotFound),
					Message: "Pet not found",
				},
			}
		}
		return &api.ErrorStatusCode{
			StatusCode: http.StatusInternalServerError,
			Response: api.Error{
				Code:    int32(http.StatusInternalServerError),
				Message: "Internal server error",
			},
		}
	}

	return nil
}
