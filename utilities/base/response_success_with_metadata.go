package base

import "backend-mental-guardians/entities"

type BaseMetadataSuccessResponse struct {
	Status   bool               `json:"status"`
	Message  string             `json:"message"`
	Metadata *entities.Metadata `json:"metadata"`
	Data     any                `json:"data"`
}

func NewMetadataSuccessResponse(message string, metadata *entities.Metadata, data any) *BaseMetadataSuccessResponse {
	return &BaseMetadataSuccessResponse{
		Status:   true,
		Message:  message,
		Metadata: metadata,
		Data:     data,
	}
}
