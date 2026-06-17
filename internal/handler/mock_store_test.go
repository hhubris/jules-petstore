package handler

import (
	"context"
	"sort"

	"petstore/internal/domain"
	"petstore/internal/storage"
)

type mockStore struct {
	pets   map[int64]*domain.Pet
	nextID int64
}

func newMockStore() domain.PetStore {
	return &mockStore{
		pets:   make(map[int64]*domain.Pet),
		nextID: 1,
	}
}

func (m *mockStore) ProviderName() string { return "mock" }

func (m *mockStore) AddPet(ctx context.Context, input domain.AddPetInput) (domain.AddPetOutput, error) {
	pet := &domain.Pet{
		ID:   m.nextID,
		Name: input.Name,
		Tag:  input.Tag,
	}
	m.pets[m.nextID] = pet
	m.nextID++
	return domain.AddPetOutput{Pet: *pet}, nil
}

func (m *mockStore) DeletePet(ctx context.Context, input domain.DeletePetInput) (domain.DeletePetOutput, error) {
	if _, ok := m.pets[input.ID]; !ok {
		return domain.DeletePetOutput{}, storage.ErrNotFound
	}
	delete(m.pets, input.ID)
	return domain.DeletePetOutput{}, nil
}

func (m *mockStore) FindPetByID(ctx context.Context, input domain.FindPetByIDInput) (domain.FindPetByIDOutput, error) {
	pet, ok := m.pets[input.ID]
	if !ok {
		return domain.FindPetByIDOutput{}, storage.ErrNotFound
	}
	return domain.FindPetByIDOutput{Pet: *pet}, nil
}

func (m *mockStore) FindPets(ctx context.Context, input domain.FindPetsInput) (domain.FindPetsOutput, error) {
	var result []domain.Pet
	for _, pet := range m.pets {
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
