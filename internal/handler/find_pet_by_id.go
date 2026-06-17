package handler

import (
	"context"
	"errors"
	"net/http"

	"petstore/internal/api"
	"petstore/internal/domain"
	"petstore/internal/storage"
)

func (s *PetStoreHandler) FindPetByID(ctx context.Context, params api.FindPetByIDParams) (*api.Pet, error) {
	input := domain.FindPetByIDInput{
		ID: params.ID,
	}

	out, err := s.store.FindPetByID(ctx, input)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, &api.ErrorStatusCode{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Code:    int32(http.StatusNotFound),
					Message: "Pet not found",
				},
			}
		}
		return nil, &api.ErrorStatusCode{
			StatusCode: http.StatusInternalServerError,
			Response: api.Error{
				Code:    int32(http.StatusInternalServerError),
				Message: "Internal server error",
			},
		}
	}

	res := &api.Pet{
		ID:   out.Pet.ID,
		Name: out.Pet.Name,
	}
	if out.Pet.Tag != nil {
		res.Tag = api.NewOptString(*out.Pet.Tag)
	}

	return res, nil
}
