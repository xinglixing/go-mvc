package app

import (
	"net/http"
	"xinglixing/go-mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
}
