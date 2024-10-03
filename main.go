package main

import (
	"backend-mental-guardians/configs"
	contentControllers "backend-mental-guardians/controllers/content"
	musicControllers "backend-mental-guardians/controllers/music"
	therapistControllers "backend-mental-guardians/controllers/therapist"
	userControllers "backend-mental-guardians/controllers/user"
	"backend-mental-guardians/repositories/mysql"
	contentRepositories "backend-mental-guardians/repositories/mysql/content"
	musicRepositories "backend-mental-guardians/repositories/mysql/music"
	therapistRepositories "backend-mental-guardians/repositories/mysql/therapist"
	userRepositories "backend-mental-guardians/repositories/mysql/user"
	"backend-mental-guardians/routes"
	contentUseCases "backend-mental-guardians/usecases/content"
	musicUseCases "backend-mental-guardians/usecases/music"
	therapistUseCases "backend-mental-guardians/usecases/therapist"
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

	therapistRepo := therapistRepositories.NewTherapistRepo(db)
	therapistUC := therapistUseCases.NewTherapistUseCase(therapistRepo)
	therapistCont := therapistControllers.NewTherapistController(therapistUC)

	route := routes.NewRouteController(userCont, musicCont, contentCont, therapistCont)
	e := echo.New()
	route.Route(e)
	e.Logger.Fatal(e.Start(":8080"))
}