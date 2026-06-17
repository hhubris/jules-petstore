package handler

import (
	"context"
	"testing"

	"petstore/internal/api"
)

func TestPetStoreHandler_AddPet(t *testing.T) {
	h := New()
	ctx := context.Background()

	req := &api.NewPet{
		Name: "Fluffy",
		Tag:  api.NewOptString("cat"),
	}

	pet, err := h.AddPet(ctx, req)
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
