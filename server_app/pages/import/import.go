package _import

import (
	"html/template"
	"log"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/import", ImportHandler)
}

func ImportHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("www/pages/import/import.partial.html")
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	data := struct {
		Title string
	}{
		Title: "Imports Page",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Could not run page", http.StatusInternalServerError)
		log.Fatalf(err.Error())
	}
}
