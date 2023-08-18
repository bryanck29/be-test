package init

import (
	"github.com/bryanck29/be-test/internal/config"
	"github.com/bryanck29/be-test/internal/db"
	"github.com/bryanck29/be-test/internal/http/controller"
	"github.com/bryanck29/be-test/internal/http/routes"
	"github.com/bryanck29/be-test/internal/repository"
	"github.com/bryanck29/be-test/internal/usecase"
	extUtils "github.com/bryanck29/be-test/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitApp() *echo.Echo {
	// Load env(s)
	extUtils.LoadEnv()

	// Init validator
	extUtils.InitValidator()

	// Init DB
	db.InitDB()

	// Init new echo instance
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// Init repositories
	repositories := repository.InitRepositories(config.Core.DB)

	// Init usecases
	usecases := usecase.InitUsecases(repositories)

	// Init controllers
	controllers := controller.InitControllers(e, usecases)

	// Init routes
	routes.InitRoutes(e, controllers)

	return e
}
