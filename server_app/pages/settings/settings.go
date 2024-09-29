package settings

import (
	"html/template"
	"log"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/settings", SettingsHandler)
}

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("www/pages/settings/settings.partial.html")
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	data := struct {
		Title         string
		StorageFolder string
		LogsFolder    string
	}{
		Title:         "Settings Page",
		StorageFolder: "c:\\go_next_funds\\storage",
		LogsFolder:    "c:\\go_next_funds\\logs",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Could not run page", http.StatusInternalServerError)
		log.Fatalf(err.Error())
	}
}
