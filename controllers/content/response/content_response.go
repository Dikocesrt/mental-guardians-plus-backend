package response

type ContentResponse struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Content      string `json:"content"`
	Category     string `json:"category"`
	Type         string `json:"type"`
	ThumbnailURL string `json:"thumbnailURL"`
	CreatedAt    string `json:"createdAt"`
}