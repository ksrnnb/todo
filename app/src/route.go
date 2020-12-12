package main

import (
	"controllers"
	"middleware"
	"net/http"
)

// パスに応じて振り分け
func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handleRoot(w, r)
	case "/favicon.ico":
		// nothing to do
	default:
		handleTodo(w, r)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controllers.Index(w, r)
	case "POST":
		controllers.CreateTodo(middleware.Csrf(w, r))
	}
}

// todo page
func handleTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controllers.ShowItem(w, r)
	case "POST":
		controllers.CreateItem(middleware.Csrf(w, r))
	}
}
