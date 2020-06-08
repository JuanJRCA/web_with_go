package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir    string = "views/layout/"
	LayoutExt    string = ".gohtml"
	InitTemplate string = "views/"
)

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, files ...string) *View {
	InitFiles(files)
	EndFiles(files)
	files = append(files, layoutFiles()...)
	t, error := template.ParseFiles(files...)

	if error != nil {
		panic(error)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}

}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)

}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := v.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func layoutFiles() []string {
	files, error := filepath.Glob(LayoutDir + "*" + LayoutExt)
	if error != nil {
		panic(error)
	}

	return files
}

func InitFiles(files []string) {
	for i, f := range files {
		files[i] = InitTemplate + f
	}

}

func EndFiles(files []string) {
	for i, f := range files {
		files[i] = f + LayoutExt
	}

}
