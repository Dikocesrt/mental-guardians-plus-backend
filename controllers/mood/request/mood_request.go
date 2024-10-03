package request

type MoodCreateRequest struct {
	Content string `json:"content" form:"content" validate:"required"`
}