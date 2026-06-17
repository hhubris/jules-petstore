package handler

import (
	"context"
	"testing"

	"petstore/internal/api"
)

func TestPetStoreHandler_FindPets(t *testing.T) {
	h := New()
	ctx := context.Background()

	// Add a few pets
	h.AddPet(ctx, &api.NewPet{Name: "Cat 1", Tag: api.NewOptString("cat")})
	h.AddPet(ctx, &api.NewPet{Name: "Dog 1", Tag: api.NewOptString("dog")})
	h.AddPet(ctx, &api.NewPet{Name: "Cat 2", Tag: api.NewOptString("cat")})

	// Find all pets
	pets, err := h.FindPets(ctx, api.FindPetsParams{})
	if err != nil {
		t.Fatalf("FindPets failed: %v", err)
	}
	if len(pets) != 3 {
		t.Errorf("expected 3 pets, got %d", len(pets))
	}

	// Find pets by tag "cat"
	catPets, err := h.FindPets(ctx, api.FindPetsParams{Tags: []string{"cat"}})
	if err != nil {
		t.Fatalf("FindPets failed: %v", err)
	}
	if len(catPets) != 2 {
		t.Errorf("expected 2 cat pets, got %d", len(catPets))
	}

	// Find with limit 1
	limitedPets, err := h.FindPets(ctx, api.FindPetsParams{Limit: api.NewOptInt32(1)})
	if err != nil {
		t.Fatalf("FindPets failed: %v", err)
	}
	if len(limitedPets) != 1 {
		t.Errorf("expected 1 pet, got %d", len(limitedPets))
	}
}
