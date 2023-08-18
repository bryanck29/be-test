package controller

import (
	"github.com/bryanck29/be-test/internal/contract"
	"github.com/bryanck29/be-test/internal/usecase"

	"github.com/labstack/echo/v4"
)

type AppControllers struct {
	AuthController contract.AuthController
	UserController contract.UserController
}

func InitControllers(e *echo.Echo, usecases usecase.AppUsecases) AppControllers {
	authController := newAuthController(e, usecases.AuthUsecase)
	userController := newUserController(e, usecases.UserUsecase)

	return AppControllers{
		AuthController: authController,
		UserController: userController,
	}
}
