package handler

import (
	"context"
	"net/http"
	"testing"

	"petstore/internal/api"
)

func TestPetStoreHandler_DeletePet(t *testing.T) {
	h := New()
	ctx := context.Background()

	// Add a pet
	req := &api.NewPet{
		Name: "Buddy",
	}
	addedPet, err := h.AddPet(ctx, req)
	if err != nil {
		t.Fatalf("AddPet failed: %v", err)
	}

	// Delete the pet
	err = h.DeletePet(ctx, api.DeletePetParams{ID: addedPet.ID})
	if err != nil {
		t.Fatalf("DeletePet failed: %v", err)
	}

	// Verify pet is gone
	_, err = h.FindPetByID(ctx, api.FindPetByIDParams{ID: addedPet.ID})
	if err == nil {
		t.Fatal("expected error finding deleted pet, got nil")
	}

	// Verify deleting non-existent pet
	err = h.DeletePet(ctx, api.DeletePetParams{ID: addedPet.ID})
	if err == nil {
		t.Fatal("expected error finding deleted pet, got nil")
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
