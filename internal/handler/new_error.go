package handler

import (
	"context"
	"net/http"

	"petstore/internal/api"
)

func (s *PetStoreHandler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	if apiErr, ok := err.(*api.ErrorStatusCode); ok {
		return apiErr
	}

	return &api.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		},
	}
}
