package main

import (
	"fmt"
	"net/http"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Primera web </h1>")
}

func main() {
	http.HandleFunc("/", HandleFunc)
	http.ListenAndServe(":8080", nil)
}
