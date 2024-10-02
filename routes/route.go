package routes

import (
	"backend-mental-guardians/controllers/user"
	myMiddlewares "backend-mental-guardians/middlewares"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	userController *user.UserController
}

func NewRouteController(userController *user.UserController) *RouteController {
	return &RouteController{
		userController: userController,
	}
}

func (routeController *RouteController) Route(e *echo.Echo) {
	myMiddlewares.LogMiddleware(e)

	userAuth := e.Group("/v1")
	userAuth.POST("/register", routeController.userController.Register)
	userAuth.POST("/login", routeController.userController.Login)
}