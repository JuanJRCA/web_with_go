package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/web_go/views"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	must(homeView.Render(w, nil))

}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	must(contactView.Render(w, nil))

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

	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":8080", r)

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
