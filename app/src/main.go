package main

import (
	"models"
	"net/http"
	"redis"
)

func init() {
	redis.Initialize()
	models.Initialize()
}

func main() {
	server := http.Server{
		Addr: ":8000",
	}

	http.HandleFunc("/", handleRequest)
	server.ListenAndServe()
}
