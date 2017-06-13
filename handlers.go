package main

import (
	"html/template"
	"log"
	"net/http"
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
	url := "https://api.darksky.net/forecast/dd7aee29471de7467a81eb91c6be98d9/28.407192799999997,-16.567539999999997?units=auto"
	weather, err := getWeather(url)
	if err != nil {
		log.Println(err)
	}

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
