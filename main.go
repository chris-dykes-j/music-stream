package main

import (
	"html/template"
	"net/http"
)

func main() {
    http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("static/dist"))))
    http.Handle("/music/", http.StripPrefix("/music/", http.FileServer(http.Dir("music"))))

    http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.tmpl.html"))
    err := tmpl.Execute(w, nil)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

