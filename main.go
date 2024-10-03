package main

import (
	"backend-mental-guardians/configs"
	contentControllers "backend-mental-guardians/controllers/content"
	moodControllers "backend-mental-guardians/controllers/mood"
	musicControllers "backend-mental-guardians/controllers/music"
	therapistControllers "backend-mental-guardians/controllers/therapist"
	userControllers "backend-mental-guardians/controllers/user"
	videoControllers "backend-mental-guardians/controllers/video"
	"backend-mental-guardians/repositories/mysql"
	contentRepositories "backend-mental-guardians/repositories/mysql/content"
	moodRepositories "backend-mental-guardians/repositories/mysql/mood"
	musicRepositories "backend-mental-guardians/repositories/mysql/music"
	therapistRepositories "backend-mental-guardians/repositories/mysql/therapist"
	userRepositories "backend-mental-guardians/repositories/mysql/user"
	videoRepositories "backend-mental-guardians/repositories/mysql/video"
	"backend-mental-guardians/routes"
	"backend-mental-guardians/usecases/chatbot"
	contentUseCases "backend-mental-guardians/usecases/content"
	moodUseCases "backend-mental-guardians/usecases/mood"
	musicUseCases "backend-mental-guardians/usecases/music"
	therapistUseCases "backend-mental-guardians/usecases/therapist"
	userUseCases "backend-mental-guardians/usecases/user"
	videoUseCases "backend-mental-guardians/usecases/video"

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

	chatbotUC := chatbot.NewChatbotUsecase()

	moodRepo := moodRepositories.NewMoodRepo(db)
	moodUC := moodUseCases.NewMoodUseCase(moodRepo)
	moodCont := moodControllers.NewMoodController(moodUC, chatbotUC)

	videoRepo := videoRepositories.NewVideoRepo(db)
	videoUC := videoUseCases.NewVideoUseCase(videoRepo)
	videoCont := videoControllers.NewVideoController(videoUC)

	route := routes.NewRouteController(userCont, musicCont, contentCont, therapistCont, moodCont, videoCont)
	e := echo.New()
	route.Route(e)
	e.Logger.Fatal(e.Start(":8080"))
}