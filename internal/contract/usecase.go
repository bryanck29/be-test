package contract

import (
	"github.com/bryanck29/be-test/internal/schema/model"
	"github.com/bryanck29/be-test/internal/schema/request"
	"github.com/bryanck29/be-test/internal/schema/response"

	"github.com/labstack/echo/v4"
)

type AuthUsecase interface {
	Login(ctx echo.Context, req request.PostLogin) (result response.PostLogin, err error)
	RefreshLogin(ctx echo.Context, req request.PostRefreshLogin) (result response.PostRefreshLogin, err error)
}

type UserUsecase interface {
	InsertUser(ctx echo.Context, req request.PostInsertUser) error
	GetUser(ctx echo.Context, req request.GetUser) (model.User, error)
	GetUsers(ctx echo.Context) ([]model.User, error)
	DeleteUser(ctx echo.Context, req request.DeleteUser) error
	UpdateUser(ctx echo.Context, req request.PutUser) (result model.User, err error)
}
