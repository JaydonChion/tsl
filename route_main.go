package main

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "indexlayout", "index", "loginForm")

}
