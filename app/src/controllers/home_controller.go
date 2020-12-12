package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"models"
	"session"

	"github.com/google/uuid"
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
	uuid := uuid.New().String()
	todo := models.Todo{UUID: uuid}
	models.Db.Create(&todo)

	http.Redirect(w, r, "/"+uuid, http.StatusMovedPermanently)
}
