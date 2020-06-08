package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/web_go/controllers"
	"github.com/web_go/views"
)

var (
	faqView *views.View
)

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	must(faqView.Render(w, nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Pagina no encontrada </h1>")
}

func main() {

	staticC := controllers.NewStatic()

	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	usersC := controllers.NewUser()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(notFound)

	r.HandleFunc("/faq", faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":8080", r)

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
