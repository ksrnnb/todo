package main

import (
	"fmt"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "top page")
}
