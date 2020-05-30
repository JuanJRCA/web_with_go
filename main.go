package main

import (
	"fmt"
	"net/http"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Vamos a jugar al Mario </h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<h1>Página contacto</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Pagina no encontrada")
	}
}

func main() {

	http.HandleFunc("/", HandleFunc)
	http.ListenAndServe(":8080", nil)
}
