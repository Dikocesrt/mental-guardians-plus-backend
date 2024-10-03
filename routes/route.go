package routes

import (
	content "backend-mental-guardians/controllers/content"
	"backend-mental-guardians/controllers/mood"
	"backend-mental-guardians/controllers/music"
	"backend-mental-guardians/controllers/therapist"
	"backend-mental-guardians/controllers/user"
	myMiddlewares "backend-mental-guardians/middlewares"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	userController *user.UserController
	musicController *music.MusicController
	contentController *content.ContentController
	therapistController *therapist.TherapistController
	moodController *mood.MoodController
}

func NewRouteController(userController *user.UserController, musicController *music.MusicController, contentController *content.ContentController, therapistController *therapist.TherapistController, moodController *mood.MoodController) *RouteController {
	return &RouteController{
		userController: userController,
		musicController: musicController,
		contentController: contentController,
		therapistController: therapistController,
		moodController: moodController,
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

	userRoute.GET("stories", routeController.contentController.GetAllStories)
	userRoute.GET("stories/:id", routeController.contentController.GetByID)

	userRoute.GET("articles", routeController.contentController.GetAllArticles)
	userRoute.GET("articles/:id", routeController.contentController.GetByID)

	userRoute.GET("profile", routeController.userController.GetProfileByID)

	userRoute.GET("therapists", routeController.therapistController.GetAll)
	userRoute.GET("therapists/:id", routeController.therapistController.GetByID)

	userRoute.POST("moods", routeController.moodController.Create)
	userRoute.GET("moods", routeController.moodController.GetAllByUserID)
}