package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
    http.Handle("/music/", http.StripPrefix("/music/", http.FileServer(http.Dir("music"))))
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

func getAllTracks() []string {
    var result []string
    files, err := os.ReadDir("./music")
    if err != nil {
        fmt.Println(err)
    }

    for _, file := range files {
        result = append(result, file.Name())
    }
    
    return result 
}
