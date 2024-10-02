package main

import (
	"backend-mental-guardians/configs"
	userControllers "backend-mental-guardians/controllers/user"
	"backend-mental-guardians/repositories/mysql"
	userRepositories "backend-mental-guardians/repositories/mysql/user"
	"backend-mental-guardians/routes"
	userUseCases "backend-mental-guardians/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.ConnectDB(configs.InitConfigMySQL())

	userRepo := userRepositories.NewUserRepo(db)
	userUC := userUseCases.NewUserUseCase(userRepo)
	userCont := userControllers.NewUserController(userUC)

	route := routes.NewRouteController(userCont)
	e := echo.New()
	route.Route(e)
	e.Logger.Fatal(e.Start(":8080"))
}