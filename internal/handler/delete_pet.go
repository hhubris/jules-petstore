package handler

import (
	"context"
	"errors"
	"net/http"

	"petstore/internal/api"
	"petstore/internal/domain"
	"petstore/internal/storage"
)

type deletePetStore interface {
	DeletePet(ctx context.Context, input domain.DeletePetInput) (domain.DeletePetOutput, error)
}

func doDeletePet(ctx context.Context, params api.DeletePetParams, store deletePetStore) error {
	input := domain.DeletePetInput{
		ID: params.ID,
	}

	_, err := store.DeletePet(ctx, input)
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

func (s *PetStoreHandler) DeletePet(ctx context.Context, params api.DeletePetParams) error {
	return doDeletePet(ctx, params, s.store)
}
