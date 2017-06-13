package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("").ParseFiles("templates/index.html", "templates/layout.html")
	err = tpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func OverviewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lat := vars["lat"]
	lon := vars["lon"]

	weather, err := getWeather(lat, lon)
	if err != nil {
		log.Println(err)
	}

	name, err := reverseGeocode(lat, lon)
	log.Println("*** NAME: " + name + " ***")

	tpl, err := template.New("").ParseFiles("templates/overview.html", "templates/layout.html")
	err = tpl.ExecuteTemplate(w, "layout", weather)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("").ParseFiles("templates/about.html", "templates/layout.html")
	err = tpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
