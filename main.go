package main

import (
	"net/http"
	"xinglixing/go-mvc/app"
)

func main() {
	app.StartApp()

	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}
