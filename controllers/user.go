package user

import (
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
