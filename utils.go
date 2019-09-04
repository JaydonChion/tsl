package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func generateHTML(w http.ResponseWriter, data interface{}, layout string, filenames ...string) {
	var files []string
	for _, file := range filenames {
		// files = append(files, fmt.Sprintf("athena/templates/%s.html", file))
		files = append(files, fmt.Sprintf("templates/%s.html", file))

	}

	templates := template.Must(template.ParseFiles(files...))

	templates.ExecuteTemplate(w, layout, data)
}
