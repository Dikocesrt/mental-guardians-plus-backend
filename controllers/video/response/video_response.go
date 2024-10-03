package response

type VideoResponse struct {
	ID      string `json:"id"`
	VideoID string `json:"videoID"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Views   int    `json:"views"`
	Likes   int    `json:"likes"`
	Labels  string `json:"labels"`
}