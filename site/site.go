package site

import (
	"os"
	"text/template"

	"github.com/morgulbrut/colorlog"
)

type Site struct {
	Name    string
	Legend  string
	Div     string
	Content string
}

func GenerateSite(path string, s Site) {
	f, err := os.Create(path)
	if err != nil {
		colorlog.Error("%s", err.Error())
		return
	}
	colorlog.Info("Executing template")
	tmpl := template.Must(template.ParseFiles("template.html"))
	err = tmpl.Execute(f, s)
	if err != nil {
		colorlog.Error("%s", err.Error())
		return
	}
}
