package video

import (
	"backend-mental-guardians/controllers/video/response"
	"backend-mental-guardians/entities/video"
	"backend-mental-guardians/utilities"
	"backend-mental-guardians/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VideoController struct {
	videoUseCase video.UseCaseInterface
}

func NewVideoController(videoUseCase video.UseCaseInterface) *VideoController {
	return &VideoController{
		videoUseCase: videoUseCase,
	}
}

func (vc *VideoController) GetAll(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	category := c.QueryParam("category")
	category = utilities.CapitalizeFirstLetter(category)

	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	_, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	videos, err := vc.videoUseCase.GetAll(*metadata, category)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	videoResponses := make([]response.VideoResponse, len(videos))
	for i, video := range videos {
		videoResponses[i] = response.VideoResponse{
			ID:        video.ID,
			VideoID:   video.VideoID,
			Title:     video.Title,
			Author:    video.Author,
			Views:     video.Views,
			Likes:     video.Likes,
			Labels:    video.Labels,
		}
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get All Video", metadata, videoResponses))
}