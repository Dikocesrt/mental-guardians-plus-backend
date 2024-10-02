package response

type MusicResponse struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Singer       string `json:"singer"`
	MusicURL     string `json:"musicURL"`
	ThumbnailURL string `json:"thumbnailURL"`
}