package story

import (
	"backend-mental-guardians/controllers/content/response"
	contentEntities "backend-mental-guardians/entities/content"
	"backend-mental-guardians/utilities"
	"backend-mental-guardians/utilities/base"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ContentController struct {
	contentUseCase contentEntities.UseCaseInterface
}

func NewContentController(contentUseCase contentEntities.UseCaseInterface) *ContentController {
	return &ContentController{
		contentUseCase: contentUseCase,
	}
}

func (s *ContentController) GetAllStories(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	category := c.QueryParam("category")

	token := c.Request().Header.Get("Authorization")
	_, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	storyData, err := s.contentUseCase.GetAll(*metadata, category, "story")
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	storyResps := make([]response.ContentResponse, len(storyData))
	for i, story := range storyData {
		storyResps[i] = response.ContentResponse{
			ID:        story.ID,
			Title:     story.Title,
			Author:    story.Author,
			Content:   story.Content,
			ThumbnailURL: story.ThumbnailURL,
			Category:  story.Category,
			Type:      story.Type,
			CreatedAt: story.CreatedAt,
		}
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get All Story", metadata, storyResps))
}

func (s *ContentController) GetAllArticles(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	category := c.QueryParam("category")

	token := c.Request().Header.Get("Authorization")
	_, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	articleData, err := s.contentUseCase.GetAll(*metadata, category, "article")
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	articleResps := make([]response.ContentResponse, len(articleData))
	for i, article := range articleData {
		articleResps[i] = response.ContentResponse{
			ID:        article.ID,
			Title:     article.Title,
			Author:    article.Author,
			Content:   article.Content,
			ThumbnailURL: article.ThumbnailURL,
			Category:  article.Category,
			Type:      article.Type,
			CreatedAt: article.CreatedAt,
		}
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get All Article", metadata, articleResps))
}

func (s *ContentController) GetByID(c echo.Context) error {
	strID := c.Param("id")
	id, _ := strconv.Atoi(strID)

	token := c.Request().Header.Get("Authorization")
	_, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	story, err := s.contentUseCase.GetByID(uint(id))
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	storyResp := response.ContentResponse{
		ID:        story.ID,
		Title:     story.Title,
		Author:    story.Author,
		Content:   story.Content,
		ThumbnailURL: story.ThumbnailURL,
		Category:  story.Category,
		Type:      story.Type,
		CreatedAt: story.CreatedAt,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Story", storyResp))
}