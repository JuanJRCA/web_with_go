package views

import (
	"html/template"
)

type View struct {
	Template *template.Template
}

func NewView(files ...string) *View {
	files = append(files, "views/layout/footer.gohtml")
	t, error := template.ParseFiles(files...)

	if error != nil {
		panic(error)
	}

	return &View{Template: t}

}
