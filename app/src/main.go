package main

import (
	"models"
	"net/http"

)

func init() {
	models.Initialize()
}

func main() {
	server := http.Server{
		Addr: ":8000",
	}

	http.HandleFunc("/", handleRequest)
	server.ListenAndServe()
}
