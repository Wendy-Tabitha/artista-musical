package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func ServeData(w http.ResponseWriter, r *http.Request) {
	var data string
	tmpl, err := template.ParseFiles("./templates/template.html")
	if err != nil {
		log.Fatal("Error parsing template")
	}

	tmpl.ExecuteTemplate(w, "template.html", data)
}