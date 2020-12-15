package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"models"
	"session"

)

// Index is top page action shows top page
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println(err)
	}

	token := session.Start(w, r)
	tmpl.Execute(w, token)
}

// CreateTodo creates Todo and shows todo page
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := models.CreateNewTodo()
	http.Redirect(w, r, "/"+todo.UUID, http.StatusMovedPermanently)
}
