package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	dir = dir + "/Music/Vocaloid"

	http.Handle("/music/", http.StripPrefix("/music/", http.FileServer(http.Dir(dir))))
	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.tmpl.html"))
	tracks := getAllTracks()
	err := tmpl.Execute(w, tracks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
