package main

import (
	"backend-mental-guardians/configs"
	musicControllers "backend-mental-guardians/controllers/music"
	userControllers "backend-mental-guardians/controllers/user"
	"backend-mental-guardians/repositories/mysql"
	musicRepositories "backend-mental-guardians/repositories/mysql/music"
	userRepositories "backend-mental-guardians/repositories/mysql/user"
	"backend-mental-guardians/routes"
	musicUseCases "backend-mental-guardians/usecases/music"
	userUseCases "backend-mental-guardians/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.ConnectDB(configs.InitConfigMySQL())

	userRepo := userRepositories.NewUserRepo(db)
	userUC := userUseCases.NewUserUseCase(userRepo)
	userCont := userControllers.NewUserController(userUC)

	musicRepo := musicRepositories.NewMusicRepo(db)
	musicUC := musicUseCases.NewMusicUseCase(musicRepo)
	musicCont := musicControllers.NewMusicController(musicUC)

	route := routes.NewRouteController(userCont, musicCont)
	e := echo.New()
	route.Route(e)
	e.Logger.Fatal(e.Start(":8080"))
}