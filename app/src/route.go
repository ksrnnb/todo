package main

import (
	"net/http"
	"controllers"
)

// パスに応じて振り分け
func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		handleRoot(w, r)
	} else {
		handleTodo(w, r)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controllers.Index(w, r)
	case "POST":
		controllers.CreateTodo(w, r)
	}
}

// todo page
func handleTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controllers.ShowItem(w, r)
	case "POST":
		controllers.CreateItem(w, r)
	}
}