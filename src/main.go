package main

import (
	"fmt"
	"github.com/codeanythingpossible/GoTimelines/timelines"
	"html/template"
	"net/http"
)

type Data struct {
	Message string
}

func main() {
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

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/update", updateHandler)
	err = http.ListenAndServe(":8080", nil)
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
