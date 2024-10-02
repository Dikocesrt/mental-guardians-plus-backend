package music

import (
	"backend-mental-guardians/controllers/music/response"
	musicEntities "backend-mental-guardians/entities/music"
	"backend-mental-guardians/utilities"
	"backend-mental-guardians/utilities/base"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MusicController struct {
	musicUseCase musicEntities.UseCaseInterface
}

func NewMusicController(musicUseCase musicEntities.UseCaseInterface) *MusicController {
	return &MusicController{
		musicUseCase: musicUseCase,
	}
}

func (musicController *MusicController) GetAll(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	token := c.Request().Header.Get("Authorization")
	_, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	musicData, err := musicController.musicUseCase.GetAll(*metadata)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	musicResps := make([]response.MusicResponse, len(musicData))

	for i, music := range musicData {
		musicResps[i] = response.MusicResponse{
			ID:           music.ID,
			Title:        music.Title,
			Singer:       music.Singer,
			MusicURL:     music.MusicURL,
			ThumbnailURL: music.ThumbnailURL,
		}
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get All Music", metadata, musicResps))
}

func (musicController *MusicController) GetByID(c echo.Context) error {
	strID := c.Param("id")
	id, _ := strconv.Atoi(strID)

	token := c.Request().Header.Get("Authorization")
	_, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	music, err := musicController.musicUseCase.GetByID(uint(id))
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	musicResp := response.MusicResponse{
		ID:           music.ID,
		Title:        music.Title,
		Singer:       music.Singer,
		MusicURL:     music.MusicURL,
		ThumbnailURL: music.ThumbnailURL,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Music", musicResp))
}