package settings

import (
	"GoNextFund/services"
	"html/template"
	"log"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/settings", SettingsHandler)
}

func GetSettings() *services.Settings {
	settingsFilepath := services.GetDefaultSettingsFilepath()
	if !services.CheckIfSettingsFileExists(settingsFilepath) {
		settings := services.DefaultSettings()
		if err := settings.StoreInFile(settingsFilepath); err != nil {
			log.Println(err)
		}
		return settings
	}
	readSettings, err := services.LoadFromFile(settingsFilepath)
	if err != nil {
		log.Println(err)
		return services.DefaultSettings()
	}
	return readSettings
}

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("www/pages/settings/settings.partial.html")
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	settings := GetSettings()

	data := struct {
		Title         string
		SettingsFile  string
		StorageFolder string
		LogsFolder    string
	}{
		Title:         "Settings Page",
		SettingsFile:  services.GetDefaultSettingsFilepath(),
		StorageFolder: settings.StorageFolder,
		LogsFolder:    settings.LogsFolder,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Could not run page", http.StatusInternalServerError)
		log.Fatalf(err.Error())
	}
}
