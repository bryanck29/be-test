package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bryanck29/be-test/internal/constant"
	"github.com/bryanck29/be-test/pkg/utils"

	"github.com/go-playground/validator"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

// SuccessResponse returns http success response
func SuccessResponse(ctx echo.Context, httpStatus int, message string, data interface{}) error {
	return utils.SuccessResponse(ctx, httpStatus, message, data)
}

/*
ErrorResponse returns http error response
The message parameter is used when you want to use custom message,
otherwise pass on an empty string to use original error message
*/
func ErrorResponse(ctx echo.Context, err error, message string, data interface{}) error {
	httpStatus, compossedMessage := LogAndComposeErrorResponse(err)
	if message == "" {
		message = compossedMessage
	}
	utils.LogError("http response: " + message)

	return utils.ErrorResponse(ctx, httpStatus, message, data)
}

// LogAndComposeErrorResponse is an helper function that helps logging and compose error response
func LogAndComposeErrorResponse(err error) (int, string) {
	utils.LogError(err.Error())
	httpStatusCode, err := ErrorMapping(err)
	message := err.Error()
	if ve, ok := err.(validator.ValidationErrors); ok {
		httpStatusCode = http.StatusBadRequest
		errFields := []string{}
		for _, fe := range ve {
			errFields = append(errFields, fe.Field())
		}
		if len(errFields) > 0 {
			message = fmt.Sprintf("invalid input for: %s", strings.Join(errFields, ", "))
		}
	} else if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == 1062 {
			httpStatusCode = http.StatusConflict
			message = constant.ErrDataExists.Error()
		}
	}

	return httpStatusCode, message
}
