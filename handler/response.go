package handler

import ()

type ErrorResponse struct {
	Detail	string
}

func NewErrorResponse(err string) *ErrorResponse {
	return &ErrorResponse{
		Detail: err,
	}
}