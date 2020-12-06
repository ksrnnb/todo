package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"models"

	"github.com/google/uuid"
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
		index(w, r)
	case "POST":
		create(w, r)
	}
}

func handleTodo(w http.ResponseWriter, r *http.Request) {
	// e.g.: a272270a-34f7-11eb-a0cf-0242ac120003
	isUUID := isUUID(r.URL.Path)

	if isUUID {
		show(w, r)
		return
	}

	error(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println(err)
	}

	uuid := uuid.New().String()

	tmpl.Execute(w, uuid)
}

// todo listを作成し、todoページを表示する
// validationは別に用意した方がいい？
func create(w http.ResponseWriter, r *http.Request) {
	uuid := r.PostFormValue("uuid")
	if isUUID(uuid) {
		todo := models.Todo{UUID: uuid}
		db.Create(&todo)

		// token := r.PostFormValue("_token")
		http.Redirect(w, r, "/" + uuid, 301)
	} else {
		http.Error(w, "post parameter is wrong", 422)
	}
}

// todoページの表示
func show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "todo page")
}

func error(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "error")
}

// パラメータがuuidかどうかを判定
func isUUID(path string) bool {
	action := strings.TrimLeft(path, "/")

	_, err := uuid.Parse(action)

	return err == nil
}
