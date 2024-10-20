package app

import (
	"fmt"
	"html/template"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.gohtml")

	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, nil)
}
