package routes

import (
	"github.com/bryanck29/be-test/internal/contract"
	"github.com/bryanck29/be-test/internal/http/middleware"

	"github.com/labstack/echo/v4"
)

func initUserRoutes(e *echo.Echo, userController contract.UserController) {
	/*
		Start of user route list
	*/
	v1 := e.Group("/v1")
	userRouterV1 := v1.Group("/user")

	// V1 routes starts here
	userRouterV1.POST("", userController.PostInsertUser, middleware.AuthAdmin)
	userRouterV1.GET("", userController.GetUsers, middleware.AuthAdmin)
	userRouterV1.GET("/:user_id", userController.GetUser, middleware.AuthUserOrAdmin)
	userRouterV1.DELETE("/:user_id", userController.DeleteUser, middleware.AuthAdmin)
	userRouterV1.PUT("/:user_id", userController.PutUser, middleware.AuthAdmin)
}
