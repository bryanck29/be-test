package utils

import (
	"net/http"
	"time"

	"github.com/bryanck29/be-test/pkg/model"

	"github.com/labstack/echo/v4"
)

// SuccessResponse returns
func SuccessResponse(ctx echo.Context, httpStatus int, message string, data interface{}) error {
	return ctx.JSON(httpStatus, model.Response{
		Status:     http.StatusText(httpStatus),
		StatusCode: httpStatus,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	})
}

// ErrorResponse returns
func ErrorResponse(ctx echo.Context, httpStatus int, message string, data interface{}) error {
	return ctx.JSON(httpStatus, model.Response{
		Status:     http.StatusText(httpStatus),
		StatusCode: httpStatus,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	})
}
