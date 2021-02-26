package services

import (
	"xinglixing/go-mvc/domain"
	"xinglixing/go-mvc/utils"
)

// GetUser Service
func GetUser(userID int64) (*domain.User, *utils.AppError) {
	return domain.GetUser(userID)
}
