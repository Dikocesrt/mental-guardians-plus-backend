package mood

import (
	"backend-mental-guardians/controllers/mood/request"
	"backend-mental-guardians/controllers/mood/response"
	chatbotEntities "backend-mental-guardians/entities/chatbot"
	moodEntities "backend-mental-guardians/entities/mood"
	"backend-mental-guardians/entities/user"
	"backend-mental-guardians/utilities"
	"backend-mental-guardians/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MoodController struct {
	moodUseCase moodEntities.UseCaseInterface
	chatbotUseCase chatbotEntities.UseCaseInterface
}

func NewMoodController(moodUseCase moodEntities.UseCaseInterface, chatbotUseCase chatbotEntities.UseCaseInterface) *MoodController {
	return &MoodController{
		moodUseCase: moodUseCase,
		chatbotUseCase: chatbotUseCase,
	}
}

func (mc *MoodController) Create(c echo.Context) error {
	var moodRequest request.MoodCreateRequest

	c.Bind(&moodRequest)

	token := c.Request().Header.Get("Authorization")
	userID, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	moodEnt := moodEntities.Mood{
		Content: moodRequest.Content,
		User : user.User{
			ID: uint(userID),
		},
	}

	result, err := mc.chatbotUseCase.GetResult(moodEnt.Content)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	
	if result == "good" || result == "Good" || result == "GOOD" || result == "good." || result == "Good." || result == "GOOD." {
		moodEnt.IsGood = true
	}else{
		moodEnt.IsGood = false
	}

	mood, err := mc.moodUseCase.Create(moodEnt)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	moodResponse := response.MoodCreateResponse{
		IsGood: mood.IsGood,
	}
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Create Mood", moodResponse))
}

func (mc *MoodController) GetAllByUserID(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	userID, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	moods, err := mc.moodUseCase.GetAllByUserID(uint(userID))
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	moodResponses := make([]response.MoodResponse, len(moods))
	for i, mood := range moods {
		moodResponses[i] = response.MoodResponse{
			ID:      mood.ID,
			Content: mood.Content,
			IsGood:  mood.IsGood,
		}
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get All Moods", moodResponses))
}