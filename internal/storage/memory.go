//go:build !pg && !mysql && !sqlite

package storage

import (
	"context"
	"sort"
	"sync"

	"petstore/internal/domain"
)

type MemoryStore struct {
	mu     sync.Mutex
	pets   map[int64]*domain.Pet
	nextID int64
}

func New() domain.PetStore {
	return &MemoryStore{
		pets:   make(map[int64]*domain.Pet),
		nextID: 1,
	}
}

func (s *MemoryStore) ProviderName() string {
	return "memory"
}

func (s *MemoryStore) AddPet(ctx context.Context, input domain.AddPetInput) (domain.AddPetOutput, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pet := &domain.Pet{
		ID:   s.nextID,
		Name: input.Name,
		Tag:  input.Tag,
	}
	s.pets[s.nextID] = pet
	s.nextID++

	return domain.AddPetOutput{Pet: *pet}, nil
}

func (s *MemoryStore) DeletePet(ctx context.Context, input domain.DeletePetInput) (domain.DeletePetOutput, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.pets[input.ID]; !ok {
		return domain.DeletePetOutput{}, ErrNotFound
	}

	delete(s.pets, input.ID)
	return domain.DeletePetOutput{}, nil
}

func (s *MemoryStore) FindPetByID(ctx context.Context, input domain.FindPetByIDInput) (domain.FindPetByIDOutput, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pet, ok := s.pets[input.ID]
	if !ok {
		return domain.FindPetByIDOutput{}, ErrNotFound
	}

	return domain.FindPetByIDOutput{Pet: *pet}, nil
}

func (s *MemoryStore) FindPets(ctx context.Context, input domain.FindPetsInput) (domain.FindPetsOutput, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []domain.Pet
	for _, pet := range s.pets {
		if input.Tags != nil {
			match := false
			for _, tag := range input.Tags {
				if pet.Tag != nil && *pet.Tag == tag {
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

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	if input.Limit != nil {
		limit := *input.Limit
		if len(result) > limit {
			result = result[:limit]
		}
	}

	return domain.FindPetsOutput{Pets: result}, nil
}
