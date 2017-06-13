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
	tpl, err := template.New("").ParseFiles("templates/overview.html", "templates/layout.html")
	err = tpl.ExecuteTemplate(w, "layout", nil)
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
