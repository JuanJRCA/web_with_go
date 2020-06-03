package views

import (
	"html/template"
)

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, files ...string) *View {
	files = append(files, "views/layout/bootstrap.gohtml",
		"views/layout/navbar.gohtml",
		"views/layout/footer.gohtml")
	t, error := template.ParseFiles(files...)

	if error != nil {
		panic(error)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}

}
