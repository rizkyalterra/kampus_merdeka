package routes

import (
	"kampus_merdeka/constants"
	"kampus_merdeka/controllers"
	"kampus_merdeka/controllers/twitters"
	"kampus_merdeka/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoute() *echo.Echo {
	e := echo.New()

	// middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.BodyDump(middlewares.Log))

	jwt := middleware.JWT([]byte(constants.SECRET_JWT))
	ev1 := e.Group("api/v1/")
	ev1.GET("users", controllers.GetUserController, jwt)
	ev1.POST("users/login", controllers.LoginController)
	ev1.POST("users/register", controllers.RegisterController)
	ev1.GET("users/:userId", controllers.DetailUserController)

	ev1.GET("tweet", twitters.GetTwitterController)
	return e
}
