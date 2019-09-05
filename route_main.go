package main

import (
	"fmt"
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

		data, err := getData("EatBook")

		if err != nil {
			panic(err)
		}

		err = t.Execute(w, data)

	case "Food":
		t := template.Must(template.ParseFiles("templates/food.html"))
		data, err := getData("Food")

		if err != nil {
			panic(err)
		}

		fmt.Println(data)

		err = t.Execute(w, data)

	case "Desserts":
		t := template.Must(template.ParseFiles("templates/desserts.html"))
		data, err := getData("Desserts")

		if err != nil {
			panic(err)
		}

		err = t.Execute(w, data)

	default:

	}

}
