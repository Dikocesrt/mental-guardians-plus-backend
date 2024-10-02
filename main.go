package main

import (
	"backend-mental-guardians/configs"
	contentControllers "backend-mental-guardians/controllers/content"
	musicControllers "backend-mental-guardians/controllers/music"
	userControllers "backend-mental-guardians/controllers/user"
	"backend-mental-guardians/repositories/mysql"
	contentRepositories "backend-mental-guardians/repositories/mysql/content"
	musicRepositories "backend-mental-guardians/repositories/mysql/music"
	userRepositories "backend-mental-guardians/repositories/mysql/user"
	"backend-mental-guardians/routes"
	contentUseCases "backend-mental-guardians/usecases/content"
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

	contentRepo := contentRepositories.NewContentRepo(db)
	contentUC := contentUseCases.NewContentUseCase(contentRepo)
	contentCont := contentControllers.NewContentController(contentUC)

	route := routes.NewRouteController(userCont, musicCont, contentCont)
	e := echo.New()
	route.Route(e)
	e.Logger.Fatal(e.Start(":8080"))
}