package main

import (
	"golearning/crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view")))

	http.Handle("/search", controller.CreateSearchResultHandler("crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8890", nil)
	if err != nil {
		panic(err)
	}
}
