package handler

import (
	"context"

	"petstore/internal/api"
)

func (s *PetStoreHandler) AddPet(ctx context.Context, req *api.NewPet) (*api.Pet, error) {
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
