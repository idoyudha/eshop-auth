package v1

import (
	"net/http"

	"github.com/google/uuid"
)

type restSuccess struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func newCreateSuccess(data any) restSuccess {
	return restSuccess{
		Code:    http.StatusCreated,
		Data:    data,
		Message: "success create",
	}
}

func newGetSuccess(data authResponse) restSuccess {
	return restSuccess{
		Code:    http.StatusOK,
		Data:    data,
		Message: "success get",
	}
}

type authResponse struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
}

func newUpdateSuccess(data any) restSuccess {
	return restSuccess{
		Code:    http.StatusOK,
		Data:    data,
		Message: "success update",
	}
}

func newDeleteSuccess() restSuccess {
	return restSuccess{
		Code:    http.StatusOK,
		Message: "success delete",
	}
}
