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
	"strconv"
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

	token := session.Start(w, r)
	uuid := helpers.GetPath(r)
	todo := models.FindTodoWithItems(uuid)

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

	data := map[string]string{
		"uuid": helpers.GetPath(r),
		"name": r.PostFormValue("name"),
	}

	models.CreateNewItem(data)

	redirectTodo(w, r)
}

// validation for create item
func validateCreateItemRequest(r *http.Request) error {
	name := r.PostFormValue("name")
	nameIsValidated, _ := regexp.MatchString(`\A[[:^cntrl:]]*\z`, name)

	if !nameIsValidated {
		return errors.New("不正な入力です")
	}

	return nil
}

// EditItemName updates item name
func EditItemName(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	uuid := helpers.GetPath(r)

	if err != nil {
		log.Println(err)
		Error(w, r)
		return
	}

	name, err := getItemName(r)

	if err != nil {
		log.Println(err)
		Error(w, r)
		return
	}

	models.UpdateItemName(id, uuid, name)

	redirectTodo(w, r)
}

func getItemName(r *http.Request) (name string, err error) {
	name = r.PostFormValue("name")
	// 制御文字以外
	reg := regexp.MustCompilePOSIX(`^[^[:cntrl:]]+$`)
	matched := reg.MatchString(name)

	if matched {
		return name, nil
	}

	err = errors.New("input error")
	return "", err
}

// HandleItemDone inverts item done
func HandleItemDone(w http.ResponseWriter, r *http.Request) {

}

// DeleteItem deletes item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// handlerでuuidは確認済
	id, err := getID(r)
	uuid := helpers.GetPath(r)

	if err != nil {
		log.Println(err)
		Error(w, r)
		return
	}

	deleteIfExists(id, uuid)

	redirectTodo(w, r)
}

// get id from input
func getID(r *http.Request) (id int, err error) {
	id, err = strconv.Atoi(r.PostFormValue("id"))
	return
}

// confirm whether todo has item or not and delete it if it exits
func deleteIfExists(id int, uuid string) {
	todo := models.FindTodoWithItems(uuid)
	item, found := todo.GetItem(id)

	if found {
		item.Delete()
	}
}

func redirectTodo(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Path, http.StatusMovedPermanently)
}

// Error shows error page
func Error(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "error")
}
