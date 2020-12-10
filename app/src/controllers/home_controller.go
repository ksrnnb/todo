package controllers

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/google/uuid"
	"models"
	"helpers"
)

// top page
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println(err)
	}

	uuid := uuid.New().String()

	tmpl.Execute(w, uuid)
}

// todo listを作成し、todoページを表示する
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	uuid := r.PostFormValue("uuid")
	if helpers.IsUUID(uuid) {
		todo := models.Todo{UUID: uuid}
		models.Db.Create(&todo)

		// token := r.PostFormValue("_token")
		http.Redirect(w, r, "/" + uuid, http.StatusMovedPermanently)
	} else {
		http.Error(w, "post parameter is wrong", http.StatusUnprocessableEntity)
	}
}