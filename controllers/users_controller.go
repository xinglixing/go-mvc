package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"xinglixing/go-mvc/services"
	"xinglixing/go-mvc/utils"
)

// GetUser handler
func GetUser(w http.ResponseWriter, r *http.Request) {
	userIDParam := r.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		responseErr := &utils.AppError{
			Message: "User ID not found",
			Status:  http.StatusNotFound,
			Code:    "not_found",
		}

		jsonVal, _ := json.Marshal(responseErr)
		w.WriteHeader(responseErr.Status)
		w.Write(jsonVal)
		return
	}

	user, appErr := services.GetUser(userID)
	if err != nil {
		jsonVal, _ := json.Marshal(appErr)
		w.WriteHeader(appErr.Status)
		w.Write([]byte(jsonVal))
		return
	}

	jsonVal, _ := json.Marshal(user)
	w.Write(jsonVal)
}
