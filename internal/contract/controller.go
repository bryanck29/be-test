package contract

import (
	"github.com/labstack/echo/v4"
)

type AuthController interface {
	PostLogin(ctx echo.Context) error
	PostRefreshLogin(ctx echo.Context) error
}

type UserController interface {
	GetUser(ctx echo.Context) error
	GetUsers(ctx echo.Context) error
	PostInsertUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
	PutUser(ctx echo.Context) error
}
