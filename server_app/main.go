package main

import (
	"GoNextFund/pages/home"
	_import "GoNextFund/pages/import"
	"GoNextFund/pages/settings"
	"fmt"
	"github.com/codeanythingpossible/GoTimelines/timelines"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Message string
}

type Navigation struct {
	PageContent template.HTML
}

func main() {
	// example of timeline use : will be removed
	timeline1, err := timelines.
		NewTimeLineBuilder[int]().
		AddMonth(2024, 1, 100).
		AddMonth(2024, 2, 200).
		AddMonth(2024, 3, 300).
		AddPeriodValue(timelines.PeriodValue[int]{
			Period: timelines.Period{
				Start: timelines.DateOnly(2024, 1, 10),
				End:   timelines.DateOnly(2024, 1, 17),
			},
			Value: 80,
		}).
		AddPeriodValue(timelines.PeriodValue[int]{
			Period: timelines.Period{
				Start: timelines.DateOnly(2024, 1, 12),
				End:   timelines.DateOnly(2024, 1, 15),
			},
			Value: 50,
		}).Build()

	if err != nil {
		fmt.Println("Erreur lors de la construction de la timeline :", err)
		return
	}

	timeline1, err = timeline1.ResolveConflicts(func(p timelines.Period, a int, b int) int {
		return a + b
	})

	for _, pv := range timeline1.Items {
		fmt.Println(pv)
	}

	// program starts here

	// configure logger
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Failed to close log file: %v", err)
		}
	}(file)

	log.SetOutput(file)

	fs := http.FileServer(http.Dir("www"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)

	home.RegisterRoutes()
	_import.RegisterRoutes()
	settings.RegisterRoutes()

	err = http.ListenAndServe(":8080", logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("%s %s %s [404]\n", r.RemoteAddr, r.Method, r.URL)
		http.NotFound(w, r)
		return
	}
	tmpl := template.Must(template.ParseFiles("www/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		println(err.Error())
		return
	}
}
