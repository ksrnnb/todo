package main

import (
	"fmt"
	"net/http"
)

func handleTop(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "top page")
}

func main() {
	server := http.Server{
		Addr: ":8000",
	}

	http.HandleFunc("/", handleTop)
	server.ListenAndServe()
}