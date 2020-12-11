package controllers

import (
	"fmt"
	"net/http"
	"html/template"
	"models"
	"helpers"
	"strings"
	"errors"
	"regexp"
	"log"
)

func displayTodoPage(w http.ResponseWriter, todo models.Todo) {
	t, _ := template.ParseFiles(
		"templates/layout.html",
		"templates/todo.html")

	t.ExecuteTemplate(w, "layout", todo)
}

// todoページの表示
func ShowItem(w http.ResponseWriter, r *http.Request) {
	// e.g.: a272270a-34f7-11eb-a0cf-0242ac120003
	err := validateShowRequest(r)

	if err != nil {
		log.Println(err)
		Error(w, r)
		return
	}

	var todo models.Todo
	path := getPath(r)
	models.Db.Where("uuid=?", path).Preload("Items").Find(&todo)

	displayTodoPage(w, todo)
}

// validation for show item
func validateShowRequest(r *http.Request) error {
	isUUID := helpers.IsUUID(getPath(r))

	if isUUID {
		return nil
	}

	return errors.New("パスがuuidではありません")
}

// itemの作成
func CreateItem(w http.ResponseWriter, r *http.Request) {
	err := validateCreateItemRequest(r)

	if err != nil {
		log.Println(err)
		Error(w, r)
		return
	}

	var todo models.Todo
	path := getPath(r)
	models.Db.Where("uuid=?", path).First(&todo)

	item := models.Item{Todo: todo, Name: r.PostFormValue("name")}
	models.Db.Create(&item)

	// Todo表示ページへ
	http.Redirect(w, r, r.URL.Path, http.StatusMovedPermanently)
}

// validation for create item
func validateCreateItemRequest(r *http.Request) error {
	r.ParseForm()
	inputs := r.PostForm

	isUUID := helpers.IsUUID(inputs["uuid"][0])

	name := inputs["name"][0]
	nameIsValidated, _ := regexp.MatchString(`\A[[:^cntrl:]]*\z`, name)

	if isUUID && nameIsValidated {
		return nil
	}

	return errors.New("不正な入力です")
}

func Error(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "error")
}

func getPath(r *http.Request) string {
	return strings.TrimLeft(r.URL.Path, "/")
}