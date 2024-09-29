package home

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/home/update", updateHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(GetHomeHtmlFilePath())
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   "Home Page",
		Message: "Hello World",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Could not run page", http.StatusInternalServerError)
		log.Fatalf(err.Error())
	}
}

func GetDefaultContent() (string, error) {
	tmpl2, err := template.ParseFiles(GetHomeHtmlFilePath())
	if err != nil {
		return "", err
	}
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Home Page",
		Message: "Hello World",
	}
	var tmp bytes.Buffer
	err = tmpl2.Execute(&tmp, data)
	if err != nil {
		return "", err
	}
	page := tmp.String()
	return page, nil
}

func GetHomeHtmlFilePath() string {
	return "www/pages/home/home.partial.html"
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("fragment").Parse(`<div hx-swap-oob="true" id="message">{{.Message}}</div>`))
	data := struct {
		Message string
	}{
		Message: "Le message a été mis à jour!",
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
