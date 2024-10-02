package response

type StoryResponse struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Content      string `json:"content"`
	Category     string `json:"category"`
	ThumbnailURL string `json:"thumbnailURL"`
}