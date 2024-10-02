package entities

type Metadata struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (metadata *Metadata) Offset() int {
	return (metadata.Page - 1) * metadata.Limit
}