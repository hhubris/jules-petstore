package domain

import (
	"context"
)

type Pet struct {
	ID   int64
	Name string
	Tag  *string
}

type AddPetInput struct {
	Name string
	Tag  *string
}

type AddPetOutput struct {
	Pet Pet
}

type DeletePetInput struct {
	ID int64
}

type DeletePetOutput struct{}

type FindPetByIDInput struct {
	ID int64
}

type FindPetByIDOutput struct {
	Pet Pet
}

type FindPetsInput struct {
	Tags  []string
	Limit *int
}

type FindPetsOutput struct {
	Pets []Pet
}

type PetStore interface {
	AddPet(ctx context.Context, input AddPetInput) (AddPetOutput, error)
	DeletePet(ctx context.Context, input DeletePetInput) (DeletePetOutput, error)
	FindPetByID(ctx context.Context, input FindPetByIDInput) (FindPetByIDOutput, error)
	FindPets(ctx context.Context, input FindPetsInput) (FindPetsOutput, error)
	ProviderName() string
}
