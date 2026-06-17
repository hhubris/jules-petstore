package handler

import (
	"context"
	"net/http"

	"petstore/internal/api"
	"petstore/internal/domain"
)

func (s *PetStoreHandler) FindPets(ctx context.Context, params api.FindPetsParams) ([]api.Pet, error) {
	input := domain.FindPetsInput{
		Tags: params.Tags,
	}
	if params.Limit.IsSet() {
		limit := int(params.Limit.Value)
		input.Limit = &limit
	}

	out, err := s.store.FindPets(ctx, input)
	if err != nil {
		return nil, &api.ErrorStatusCode{
			StatusCode: http.StatusInternalServerError,
			Response: api.Error{
				Code:    int32(http.StatusInternalServerError),
				Message: "Internal server error",
			},
		}
	}

	var result []api.Pet
	for _, p := range out.Pets {
		apiPet := api.Pet{
			ID:   p.ID,
			Name: p.Name,
		}
		if p.Tag != nil {
			apiPet.Tag = api.NewOptString(*p.Tag)
		}
		result = append(result, apiPet)
	}

	return result, nil
}
