package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8000",
	}

	http.HandleFunc("/", top)
	server.ListenAndServe()
}
