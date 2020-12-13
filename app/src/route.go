package main

import (
	"controllers"
	"helpers"
	"middleware"
	"net/http"
)

// パスに応じて振り分け
func handleRequest(w http.ResponseWriter, r *http.Request) {
	path := helpers.GetPath(r)

	switch {
	case path == "": // top page
		handleRoot(w, r)
	case helpers.IsUUID(path): // todo page
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
