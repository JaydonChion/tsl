package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "indexlayout", "index", "home")

}

func viewFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	switch pagename := vars["page"]; pagename {
	case "EatBook":
		t := template.Must(template.ParseFiles("templates/eatbook.html"))

		// jobData, err := getJobData(id)
		// if err != nil {
		// 	panic(err)
		// }

		// categories := map[int]string{
		// 	1: "Cleaning",
		// 	2: "Courier",
		// 	3: "Fixing",
		// 	4: "Care service",
		// 	5: "Sharing",
		// 	6: "Delivery",
		// 	7: "Others",
		// }

		// temp := make(map[string]interface{})
		// temp["jobData"] = jobData
		// temp["jobDataCat"] = categories[jobData.Job_cat]
		getEatbookData()

		err := t.Execute(w, nil)
		if err != nil {
			panic(err)
		}

	case "Food":
		t := template.Must(template.ParseFiles("templates/food.html"))
		err := t.Execute(w, nil)
		if err != nil {
			panic(err)
		}

	case "Desserts":
		t := template.Must(template.ParseFiles("templates/desserts.html"))
		err := t.Execute(w, nil)
		if err != nil {
			panic(err)
		}

	default:

	}

}
