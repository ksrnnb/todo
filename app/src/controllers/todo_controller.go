package controllers

import (
	"helpers"
	"html/template"
	"models"
	"net/http"
	"session"
	"strconv"
	"validations"
)

// ExecuteTemplateに渡すために構造体にまとめる
type templateParam struct {
	Token string
	Todo  models.Todo
}

func displayTodoPage(w http.ResponseWriter, params templateParam) {
	t, _ := template.ParseFiles(
		"templates/layout.html",
		"templates/todo.html")

	t.ExecuteTemplate(w, "layout", params)
}

// ShowItem shows todo page
func ShowItem(w http.ResponseWriter, r *http.Request) {
	// e.g.: a272270a-34f7-11eb-a0cf-0242ac120003
	validated := validations.ValidateShowItem(r)

	if !validated {
		redirectTodo(w, r)
		return
	}

	token := session.Start(w, r)
	uuid := helpers.GetPath(r)
	todo := models.FindTodoWithItems(uuid)

	params := templateParam{
		Token: token,
		Todo:  todo,
	}

	displayTodoPage(w, params)
}

// CreateItem creates new todo item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	validated := validations.ValidateCreateItem(r)

	if !validated {
		redirectTodo(w, r)
		return
	}

	data := map[string]string{
		"uuid": helpers.GetPath(r),
		"name": r.PostFormValue("name"),
	}

	models.CreateNewItem(data)

	redirectTodo(w, r)
}

// EditItemName updates item name
func EditItemName(w http.ResponseWriter, r *http.Request) {
	validated := validations.ValidateEditItemName(r)

	if !validated {
		redirectTodo(w, r)
		return
	}

	id, uuid, name := getID(r), getUUID(r), getItemName(r)

	models.UpdateItemName(id, uuid, name)

	redirectTodo(w, r)
}

// HandleItemDone inverts item done
func HandleItemDone(w http.ResponseWriter, r *http.Request) {
	validated := validations.ValidateHandleDone(r)

	if !validated {
		redirectTodo(w, r)
		return
	}

	id, uuid := getID(r), getUUID(r)

	models.UpdateItemDone(id, uuid)

	redirectTodo(w, r)
}

// DeleteItem deletes item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	validated := validations.ValidateDeleteItem(r)

	if !validated {
		redirectTodo(w, r)
		return
	}

	id, uuid := getID(r), getUUID(r)

	models.DeleteItem(id, uuid)

	redirectTodo(w, r)
}

// get id from input
func getID(r *http.Request) (id int) {
	id, _ = strconv.Atoi(r.PostFormValue("id"))
	return
}

// get uuid
func getUUID(r *http.Request) (uuid string) {
	uuid = helpers.GetPath(r)
	return
}

// get name
func getItemName(r *http.Request) (name string) {
	name = r.PostFormValue("name")
	return
}

// redirect to todo page
func redirectTodo(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Path, http.StatusMovedPermanently)
}
