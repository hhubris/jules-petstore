package server

import (
	"context"
	"net/http"
	"testing"

	"petstore/internal/api"
)

func TestPetStoreServer_AddPet(t *testing.T) {
	srv := NewPetStoreServer()
	ctx := context.Background()

	req := &api.NewPet{
		Name: "Fluffy",
		Tag:  api.NewOptString("cat"),
	}

	pet, err := srv.AddPet(ctx, req)
	if err != nil {
		t.Fatalf("AddPet failed: %v", err)
	}

	if pet.Name != "Fluffy" {
		t.Errorf("expected name Fluffy, got %s", pet.Name)
	}
	if pet.ID != 1 {
		t.Errorf("expected ID 1, got %d", pet.ID)
	}
}

func TestPetStoreServer_FindPetByID(t *testing.T) {
	srv := NewPetStoreServer()
	ctx := context.Background()

	// Add a pet first
	req := &api.NewPet{
		Name: "Rex",
	}
	addedPet, err := srv.AddPet(ctx, req)
	if err != nil {
		t.Fatalf("AddPet failed: %v", err)
	}

	// Find the pet
	pet, err := srv.FindPetByID(ctx, api.FindPetByIDParams{ID: addedPet.ID})
	if err != nil {
		t.Fatalf("FindPetByID failed: %v", err)
	}
	if pet.Name != "Rex" {
		t.Errorf("expected name Rex, got %s", pet.Name)
	}

	// Find a non-existent pet
	_, err = srv.FindPetByID(ctx, api.FindPetByIDParams{ID: 999})
	if err == nil {
		t.Fatal("expected error finding non-existent pet, got nil")
	}

	var errStatusCode *api.ErrorStatusCode
	switch e := err.(type) {
	case *api.ErrorStatusCode:
		errStatusCode = e
	default:
		t.Fatalf("expected *api.ErrorStatusCode, got %T", err)
	}

	if errStatusCode.StatusCode != http.StatusNotFound {
		t.Errorf("expected status code %d, got %d", http.StatusNotFound, errStatusCode.StatusCode)
	}
}

func TestPetStoreServer_DeletePet(t *testing.T) {
	srv := NewPetStoreServer()
	ctx := context.Background()

	// Add a pet
	req := &api.NewPet{
		Name: "Buddy",
	}
	addedPet, err := srv.AddPet(ctx, req)
	if err != nil {
		t.Fatalf("AddPet failed: %v", err)
	}

	// Delete the pet
	err = srv.DeletePet(ctx, api.DeletePetParams{ID: addedPet.ID})
	if err != nil {
		t.Fatalf("DeletePet failed: %v", err)
	}

	// Verify pet is gone
	_, err = srv.FindPetByID(ctx, api.FindPetByIDParams{ID: addedPet.ID})
	if err == nil {
		t.Fatal("expected error finding deleted pet, got nil")
	}
}

func TestPetStoreServer_FindPets(t *testing.T) {
	srv := NewPetStoreServer()
	ctx := context.Background()

	// Add a few pets
	srv.AddPet(ctx, &api.NewPet{Name: "Cat 1", Tag: api.NewOptString("cat")})
	srv.AddPet(ctx, &api.NewPet{Name: "Dog 1", Tag: api.NewOptString("dog")})
	srv.AddPet(ctx, &api.NewPet{Name: "Cat 2", Tag: api.NewOptString("cat")})

	// Find all pets
	pets, err := srv.FindPets(ctx, api.FindPetsParams{})
	if err != nil {
		t.Fatalf("FindPets failed: %v", err)
	}
	if len(pets) != 3 {
		t.Errorf("expected 3 pets, got %d", len(pets))
	}

	// Find pets by tag "cat"
	catPets, err := srv.FindPets(ctx, api.FindPetsParams{Tags: []string{"cat"}})
	if err != nil {
		t.Fatalf("FindPets failed: %v", err)
	}
	if len(catPets) != 2 {
		t.Errorf("expected 2 cat pets, got %d", len(catPets))
	}

	// Find with limit 1
	limitedPets, err := srv.FindPets(ctx, api.FindPetsParams{Limit: api.NewOptInt32(1)})
	if err != nil {
		t.Fatalf("FindPets failed: %v", err)
	}
	if len(limitedPets) != 1 {
		t.Errorf("expected 1 pet, got %d", len(limitedPets))
	}
}
