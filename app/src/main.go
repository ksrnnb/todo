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
	assetsHandler := http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))

	http.HandleFunc("/", handleRequest)
	http.Handle("/assets/", assetsHandler)

	server.ListenAndServe()
}
