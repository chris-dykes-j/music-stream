package main

import (
	"html/template"
	"net/http"
)

func main() {
    http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html"))
    err := tmpl.Execute(w, nil)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
