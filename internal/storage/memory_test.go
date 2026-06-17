package storage

import (
	"context"
	"testing"

	"petstore/internal/domain"
)

func TestMemoryStore_AddPet(t *testing.T) {
	s := New()
	ctx := context.Background()

	val := "dog"
	input := domain.AddPetInput{
		Name: "Fido",
		Tag:  &val,
	}

	out, err := s.AddPet(ctx, input)
	if err != nil {
		t.Fatalf("AddPet failed: %v", err)
	}

	if out.Pet.ID != 1 {
		t.Errorf("expected ID 1, got %d", out.Pet.ID)
	}
	if out.Pet.Name != "Fido" {
		t.Errorf("expected Name Fido, got %s", out.Pet.Name)
	}
	if out.Pet.Tag == nil || *out.Pet.Tag != "dog" {
		t.Errorf("expected Tag dog")
	}
}

func TestMemoryStore_FindPetByID(t *testing.T) {
	s := New()
	ctx := context.Background()

	// Add a pet first
	input := domain.AddPetInput{Name: "Rex"}
	addOut, err := s.AddPet(ctx, input)
	if err != nil {
		t.Fatalf("AddPet failed: %v", err)
	}

	// Find the pet
	findOut, err := s.FindPetByID(ctx, domain.FindPetByIDInput{ID: addOut.Pet.ID})
	if err != nil {
		t.Fatalf("FindPetByID failed: %v", err)
	}
	if findOut.Pet.Name != "Rex" {
		t.Errorf("expected Name Rex, got %s", findOut.Pet.Name)
	}

	// Try finding a non-existent pet
	_, err = s.FindPetByID(ctx, domain.FindPetByIDInput{ID: 999})
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func TestMemoryStore_DeletePet(t *testing.T) {
	s := New()
	ctx := context.Background()

	// Add a pet first
	input := domain.AddPetInput{Name: "To Delete"}
	addOut, err := s.AddPet(ctx, input)
	if err != nil {
		t.Fatalf("AddPet failed: %v", err)
	}

	// Delete the pet
	_, err = s.DeletePet(ctx, domain.DeletePetInput{ID: addOut.Pet.ID})
	if err != nil {
		t.Fatalf("DeletePet failed: %v", err)
	}

	// Try finding the deleted pet
	_, err = s.FindPetByID(ctx, domain.FindPetByIDInput{ID: addOut.Pet.ID})
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}

	// Try deleting a non-existent pet
	_, err = s.DeletePet(ctx, domain.DeletePetInput{ID: 999})
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func TestMemoryStore_FindPets(t *testing.T) {
	s := New()
	ctx := context.Background()

	// Add pets
	valCat := "cat"
	valDog := "dog"
	s.AddPet(ctx, domain.AddPetInput{Name: "Kitty", Tag: &valCat})
	s.AddPet(ctx, domain.AddPetInput{Name: "Doggo", Tag: &valDog})
	s.AddPet(ctx, domain.AddPetInput{Name: "Mittens", Tag: &valCat})

	// Find all
	outAll, err := s.FindPets(ctx, domain.FindPetsInput{})
	if err != nil {
		t.Fatalf("FindPets failed: %v", err)
	}
	if len(outAll.Pets) != 3 {
		t.Errorf("expected 3 pets, got %d", len(outAll.Pets))
	}

	// Filter by tag
	outTags, err := s.FindPets(ctx, domain.FindPetsInput{Tags: []string{"cat"}})
	if err != nil {
		t.Fatalf("FindPets failed: %v", err)
	}
	if len(outTags.Pets) != 2 {
		t.Errorf("expected 2 cats, got %d", len(outTags.Pets))
	}

	// Limit
	limit := 2
	outLimit, err := s.FindPets(ctx, domain.FindPetsInput{Limit: &limit})
	if err != nil {
		t.Fatalf("FindPets failed: %v", err)
	}
	if len(outLimit.Pets) != 2 {
		t.Errorf("expected limit 2, got %d", len(outLimit.Pets))
	}
}
