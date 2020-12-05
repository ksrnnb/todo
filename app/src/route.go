package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {

		switch r.Method {
		case "GET":
			index(w, r)
			// case: "POST"
			// 	create?
		}
	} else {
		// e.g.: a272270a-34f7-11eb-a0cf-0242ac120003
		isUuid := isUuid(r.URL.Path)

		if isUuid {
			show(w, r)
			return
		} else {
			error(w, r)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "top page")
}

func show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "todo page")
}

func error(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "error")
}

/**
　*	パスがUUIDかどうかを判定
　*	@param string path
　*	@return bool
　*/
func isUuid(path string) bool {
	action := strings.TrimLeft(path, "/")

	_, err := uuid.Parse(action)

	return err == nil
}
