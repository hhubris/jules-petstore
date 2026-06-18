package handler

import (
	"petstore/internal/api"
	"petstore/internal/domain"
)

var _ api.Handler = (*PetStoreHandler)(nil)

type PetStoreHandler struct {
	store domain.PetStore
}

func New(store domain.PetStore) *PetStoreHandler {
	return &PetStoreHandler{
		store: store,
	}
}
