package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Message string
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/update", updateHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, Data{Message: "Bienvenue dans GoNextFund!"})
	if err != nil {
		return
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("fragment").Parse(`<div hx-swap-oob="true" id="message">{{.Message}}</div>`))
	tmpl.Execute(w, Data{Message: "Le message a été mis à jour!"})
}
