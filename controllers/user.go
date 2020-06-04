package user

import (
	"fmt"
	"net/http"

	"github.com/web_go/views"
)

type User struct {
	NewView *views.View
}

func NewUser() *User {

	return &User{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

func (u *User) New(w http.ResponseWriter, r *http.Request) {
	error := u.NewView.Render(w, nil)
	if error != nil {
		panic(error)
	}
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, r.PostForm["email"])
	fmt.Fprint(w, r.PostForm["password"])
	fmt.Fprint(w, "Está haciendo un POST")
}
