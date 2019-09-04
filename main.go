package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", index)
	// router.HandleFunc("/eatbook", getEatBook)
	// router.HandleFunc("/food", getFood)
	// router.HandleFunc("/desserts", getDesserts)

	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/",
		http.FileServer(http.Dir("public/css/"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/",
		http.FileServer(http.Dir("public/js/"))))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/",
		http.FileServer(http.Dir("public/assets/"))))

	server := &http.Server{
		Addr:           ":3000",
		Handler:        router,
		ReadTimeout:    time.Duration(10 * int64(time.Second)),
		WriteTimeout:   time.Duration(600 * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

}
