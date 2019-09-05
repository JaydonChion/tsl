package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "indexlayout", "index", "home")

}

func viewFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	latitude, err := strconv.ParseFloat(vars["latitude"], 64)
	longitude, err := strconv.ParseFloat(vars["longitude"], 64)
	distance, err := strconv.Atoi(vars["distance"])
	fmt.Println(longitude)
	if err != nil {
	}

	switch pagename := vars["page"]; pagename {
	case "EatBook":
		t := template.Must(template.ParseFiles("templates/eatbook.html"))

		data, err := getData("EatBook", latitude, longitude, distance)

		if err != nil {
			panic(err)
		}

		err = t.Execute(w, data)

	case "Food":
		t := template.Must(template.ParseFiles("templates/food.html"))
		data, err := getData("Food", latitude, longitude, distance)

		if err != nil {
			panic(err)
		}

		err = t.Execute(w, data)

	case "Desserts":
		t := template.Must(template.ParseFiles("templates/desserts.html"))
		data, err := getData("Desserts", latitude, longitude, distance)

		if err != nil {
			panic(err)
		}

		err = t.Execute(w, data)

	default:

	}

}
