package routes

import (
	"github.com/bryanck29/be-test/internal/contract"
	"github.com/bryanck29/be-test/internal/http/middleware"

	"github.com/labstack/echo/v4"
)

func initAuthRoutes(e *echo.Echo, authController contract.AuthController) {
	/*
		Start of auth route list
	*/
	v1 := e.Group("/v1")
	authRouterV1 := v1.Group("/auth")

	// V1 routes starts here
	authRouterV1.POST("/login", authController.PostLogin)
	authRouterV1.POST("/refresh", authController.PostRefreshLogin, middleware.AuthUserOrAdmin)
}
