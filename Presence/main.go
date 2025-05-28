package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles("static/index.html", "static/create.html", "static/delete.html"))

type StatusMap struct {
	Name   string
	Status string
}



func main() {
	staticDir := http.Dir("./static")
	fs := http.FileServer(staticDir)

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/create/", createHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.HandleFunc("/remove/", removeHandler)
	http.HandleFunc("/update/", updateHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
