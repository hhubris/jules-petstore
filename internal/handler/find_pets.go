package handler

import (
	"context"
	"sort"

	"petstore/internal/api"
)

func (s *PetStoreHandler) FindPets(ctx context.Context, params api.FindPetsParams) ([]api.Pet, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []api.Pet
	for _, pet := range s.pets {
		if params.Tags != nil {
			// Basic filtering if tag is provided
			match := false
			for _, tag := range params.Tags {
				if pet.Tag.IsSet() && pet.Tag.Value == tag {
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

	// Sort by ID to make it deterministic
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	if params.Limit.IsSet() {
		limit := int(params.Limit.Value)
		if len(result) > limit {
			result = result[:limit]
		}
	}

	return result, nil
}
