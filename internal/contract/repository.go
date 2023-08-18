package contract

import (
	"github.com/bryanck29/be-test/internal/schema/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	InsertUser(ctx echo.Context, payload model.User) (err error)
	GetUsers(ctx echo.Context) (results []model.User, err error)
	GetUser(ctx echo.Context, userId uuid.UUID) (model.User, error)
	DeleteUser(ctx echo.Context, userId uuid.UUID) (err error)
	UpdateUser(ctx echo.Context, payload model.User) (result model.User, err error)
	GetUserByUsername(ctx echo.Context, username string) (result model.User, err error)
}
