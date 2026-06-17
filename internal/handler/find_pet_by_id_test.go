package handler

import (
	"context"
	"net/http"
	"testing"

	"petstore/internal/api"
)

func TestPetStoreHandler_FindPetByID(t *testing.T) {
	h := New()
	ctx := context.Background()

	// Add a pet first
	req := &api.NewPet{
		Name: "Rex",
	}
	addedPet, err := h.AddPet(ctx, req)
	if err != nil {
		t.Fatalf("AddPet failed: %v", err)
	}

	// Find the pet
	pet, err := h.FindPetByID(ctx, api.FindPetByIDParams{ID: addedPet.ID})
	if err != nil {
		t.Fatalf("FindPetByID failed: %v", err)
	}
	if pet.Name != "Rex" {
		t.Errorf("expected name Rex, got %s", pet.Name)
	}

	// Find a non-existent pet
	_, err = h.FindPetByID(ctx, api.FindPetByIDParams{ID: 999})
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
