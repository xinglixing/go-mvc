package domain

import (
	"fmt"
	"net/http"
	"xinglixing/go-mvc/utils"
)

var (
	users = map[int64]*User{
		1: {ID: 1, FirstName: "One", LastName: "Test", Email: "test_one@gmail.com"},
		2: {ID: 2, FirstName: "Two", LastName: "Test", Email: "test_two@gmail.com"},
		3: {ID: 3, FirstName: "John", LastName: "Doe", Email: "john.doe@test.com"},
	}
)

// GetUser get user from db
func GetUser(userID int64) (*User, *utils.AppError) {
	user := users[userID]
	if user == nil {
		return nil, &utils.AppError{
			Message: fmt.Sprintf(fmt.Sprintf("User %v was not found", userID)),
			Status:  http.StatusNotFound,
			Code:    "not_found",
		}
	}
	return user, nil
}
