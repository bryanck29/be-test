package routes

import (
	"github.com/bryanck29/be-test/internal/http/controller"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// InitRoutes initialize routes
func InitRoutes(e *echo.Echo, appController controller.AppControllers) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	initAuthRoutes(e, appController.AuthController)
	initUserRoutes(e, appController.UserController)
}
