package handler

import (
	"context"

	"petstore/internal/api"
	"petstore/internal/domain"
)

func (s *PetStoreHandler) AddPet(ctx context.Context, req *api.NewPet) (*api.Pet, error) {
	input := domain.AddPetInput{
		Name: req.Name,
	}
	if req.Tag.IsSet() {
		val := req.Tag.Value
		input.Tag = &val
	}

	out, err := s.store.AddPet(ctx, input)
	if err != nil {
		return nil, err
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
