package controllers

import (
	"errors"
	"fmt"
	"helpers"
	"html/template"
	"log"
	"models"
	"net/http"
	"regexp"
	"session"
)

// ExecuteTemplateに渡すために構造体にまとめる
type paramsStruct struct {
	Token string
	Todo  models.Todo
}

func displayTodoPage(w http.ResponseWriter, params paramsStruct) {
	t, _ := template.ParseFiles(
		"templates/layout.html",
		"templates/todo.html")

	t.ExecuteTemplate(w, "layout", params)
}

// ShowItem shows todo page
func ShowItem(w http.ResponseWriter, r *http.Request) {
	// e.g.: a272270a-34f7-11eb-a0cf-0242ac120003
	err := validateShowRequest(r)

	if err != nil {
		log.Println(err)
		Error(w, r)
		return
	}

	var todo models.Todo
	path := helpers.GetPath(r)
	models.Db.Where("uuid=?", path).Preload("Items").Find(&todo)

	token := session.Start(w, r)

	params := paramsStruct{
		Token: token,
		Todo:  todo,
	}

	displayTodoPage(w, params)
}

// validation for show item
func validateShowRequest(r *http.Request) error {
	isUUID := helpers.IsUUID(helpers.GetPath(r))

	if isUUID {
		return nil
	}

	return errors.New("パスがuuidではありません")
}

// CreateItem creates new todo item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	err := validateCreateItemRequest(r)

	if err != nil {
		log.Println(err)
		Error(w, r)
		return
	}

	var todo models.Todo
	path := helpers.GetPath(r)
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

// Error shows error page
func Error(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "error")
}
