package handler

import (
	"context"

	"petstore/internal/api"
	"petstore/internal/domain"
)

type addPetStore interface {
	AddPet(ctx context.Context, input domain.AddPetInput) (domain.AddPetOutput, error)
}

func doAddPet(ctx context.Context, req *api.NewPet, store addPetStore) (*api.Pet, error) {
	input := domain.AddPetInput{
		Name: req.Name,
	}
	if req.Tag.IsSet() {
		val := req.Tag.Value
		input.Tag = &val
	}

	out, err := store.AddPet(ctx, input)
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

func (s *PetStoreHandler) AddPet(ctx context.Context, req *api.NewPet) (*api.Pet, error) {
	return doAddPet(ctx, req, s.store)
}
