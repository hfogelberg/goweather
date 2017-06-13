package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var tpl *template.Template

func main() {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/forecast/{lat}/{lon}", OverviewHandler).Methods("GET")
	router.HandleFunc("/about", AboutHandler).Methods("GET")

	mux := http.NewServeMux()
	mux.Handle("/", router)

	static := http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
	router.PathPrefix("/public").Handler(static)

	n := negroni.Classic()
	n.UseHandler(mux)
	http.ListenAndServe(":8080", n)
}
