package handler

import (
	"sync"

	"petstore/internal/api"
)

var _ api.Handler = (*PetStoreHandler)(nil)

type PetStoreHandler struct {
	mu     sync.Mutex
	pets   map[int64]*api.Pet
	nextID int64
}

func New() *PetStoreHandler {
	return &PetStoreHandler{
		pets:   make(map[int64]*api.Pet),
		nextID: 1,
	}
}
