package controllers

import (
	"fmt"
	"net/http"
	"html/template"
	"models"
	"helpers"
	"strings"
	"errors"
	"log"
)

func parseFiles() *template.Template{
	t, _ := template.ParseFiles(
		"templates/layout.html",
		"templates/todo.html")

	return t
}

// todoページの表示
func ShowItem(w http.ResponseWriter, r *http.Request) {
	// e.g.: a272270a-34f7-11eb-a0cf-0242ac120003
	err := validateShowRequest(r)

	if err != nil {
		log.Fatalln(err)
	}

	t := parseFiles()
	action := getAction(r)

	var todo models.Todo
	models.Db.Where("uuid=?", action).First(&todo)

	t.ExecuteTemplate(w, "layout", todo)
}

// validation for show item
func validateShowRequest(r *http.Request) error {
	isUUID := helpers.IsUUID(getAction(r))

	fmt.Println(getAction(r))
	if isUUID {
		return nil
	}

	// TODO: faviconのリクエストでここに引っかかっている！
	return errors.New("パスがuuidではありません")
}

// itemの作成
func CreateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST!")
}

func Error(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "error")
}

func getAction(r *http.Request) string {
	return strings.TrimLeft(r.URL.Path, "/")
}