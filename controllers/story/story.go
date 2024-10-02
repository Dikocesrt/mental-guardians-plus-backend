package story

import (
	"backend-mental-guardians/controllers/story/response"
	storyEntities "backend-mental-guardians/entities/story"
	"backend-mental-guardians/utilities"
	"backend-mental-guardians/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StoryController struct {
	storyUseCase storyEntities.UseCaseInterface
}

func NewStoryController(storyUseCase storyEntities.UseCaseInterface) *StoryController {
	return &StoryController{
		storyUseCase: storyUseCase,
	}
}

func (s *StoryController) GetAll(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	category := c.QueryParam("category")

	token := c.Request().Header.Get("Authorization")
	_, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	storyData, err := s.storyUseCase.GetAll(*metadata, category)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	storyResps := make([]response.StoryResponse, len(storyData))
	for i, story := range storyData {
		storyResps[i] = response.StoryResponse{
			ID:        story.ID,
			Title:     story.Title,
			Author:    story.Author,
			Content:   story.Content,
			ThumbnailURL: story.ThumbnailURL,
			Category:  story.Category,
		}
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get All Story", metadata, storyResps))
}