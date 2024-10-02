package routes

import (
	"backend-mental-guardians/controllers/music"
	"backend-mental-guardians/controllers/user"
	myMiddlewares "backend-mental-guardians/middlewares"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	userController *user.UserController
	musicController *music.MusicController
}

func NewRouteController(userController *user.UserController, musicController *music.MusicController) *RouteController {
	return &RouteController{
		userController: userController,
		musicController: musicController,
	}
}

func (routeController *RouteController) Route(e *echo.Echo) {
	myMiddlewares.LogMiddleware(e)

	userAuth := e.Group("/v1")
	userAuth.POST("/register", routeController.userController.Register)
	userAuth.POST("/login", routeController.userController.Login)

	userRoute := userAuth.Group("/")
	userRoute.GET("musics", routeController.musicController.GetAll)
	userRoute.GET("musics/:id", routeController.musicController.GetByID)
}