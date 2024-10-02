package utilities

import (
	"backend-mental-guardians/entities"
	"strconv"
)

func GetMetadata(pageParam, limitParam string) *entities.Metadata {
	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}
	return &entities.Metadata{
		Page:  page,
		Limit: limit,
	}
}