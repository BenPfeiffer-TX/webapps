package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var templates = template.Must(template.ParseFiles("static/index.html"))

type StatusMap struct {
	Name   string
	Status string
}

func getStatus() ([]StatusMap, error) {
	file, err := os.Open("status.json")
	var statusArr []StatusMap
	if err != nil {
		fmt.Println("Error opening the json!" + err.Error())
		return statusArr, err
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	err = dec.Decode(&statusArr)
	if err != nil {
		fmt.Println("Error decoding the json!" + err.Error())
		return statusArr, err
	}

	return statusArr, err
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	statusArr, err := getStatus()
	data := map[string][]StatusMap{"StatusMap": statusArr}
	err = templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

}

func updateHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	staticDir := http.Dir("./static")
	fs := http.FileServer(staticDir)

	http.Handle("/static", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/create/", createHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.HandleFunc("/update/", updateHandler)

	log.Fatal(http.ListenAndServe(":8888", nil))
}
