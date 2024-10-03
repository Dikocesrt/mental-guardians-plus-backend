package response

type MoodResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	IsGood  bool   `json:"isGood"`
}