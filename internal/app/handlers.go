package app

import (
	"fmt"
	"github.com/astlaure/httpmirror/internal/messages"
	"html/template"
	"net/http"
	"strconv"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.gohtml")

	if err != nil {
		fmt.Println(err)
	}

	requests := messages.RetrieveRequests()

	tmpl.Execute(w, map[string]interface{}{
		"Requests": requests,
	})
}

func handleCompare(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/compare.gohtml")

	if err != nil {
		fmt.Println(err)
	}

	requestID := r.URL.Query().Get("request_id")
	id, err := strconv.ParseUint(requestID, 10, 32)

	mess := messages.RetrieveMessagesByRequestID(uint(id))

	tmpl.Execute(w, map[string]interface{}{
		"Active":  mess[0],
		"Preview": mess[1],
	})
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.gohtml")

	if err != nil {
		fmt.Println(err)
	}

	requests := messages.RetrieveRequests()

	tmpl.Execute(w, map[string]interface{}{
		"Requests": requests,
	})
}
