package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/web_go/controllers"
)

func main() {

	staticC := controllers.NewStatic()
	usersC := controllers.NewUser()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":8080", r)

}
