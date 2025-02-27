package entity

import (
	"github.com/susatyo441/go-ta-utils/response"
	"github.com/gofiber/fiber/v2"
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func (e *HttpError) SendResponse(ctx *fiber.Ctx) error {
	return response.SendResponse(ctx, e.Code, nil, e.Message)
}

func InternalServerError(message string) *HttpError {
	return &HttpError{
		Code:    fiber.StatusInternalServerError,
		Message: message,
	}
}

func BadRequest(message string) *HttpError {
	return &HttpError{
		Code:    fiber.StatusBadRequest,
		Message: message,
	}
}

func Unauthorized(message string) *HttpError {
	return &HttpError{
		Code:    fiber.StatusUnauthorized,
		Message: message,
	}
}

func Forbidden(message string) *HttpError {
	return &HttpError{
		Code:    fiber.StatusForbidden,
		Message: message,
	}
}

func NotFound(message string) *HttpError {
	return &HttpError{
		Code:    fiber.StatusNotFound,
		Message: message,
	}
}
