package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var (
	homeTemplate   *template.Template
	contacTemplate *template.Template
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	error := homeTemplate.Execute(w, nil)
	if error != nil {
		panic(error)
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	error := contacTemplate.Execute(w, nil)
	if error != nil {
		panic(error)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h1>Pagina de preguntas usuales </h1>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Pagina no encontrada </h1>")
}

func main() {
	var error error
	homeTemplate, error = template.ParseFiles("views/home.gohtml",
		"views/layout/footer.gohtml")
	if error != nil {
		panic(error)
	}
	contacTemplate, error = template.ParseFiles("views/contact.gohtml",
		"views/layout/footer.gohtml")
	if error != nil {
		panic(error)
	}

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":8080", r)

}
