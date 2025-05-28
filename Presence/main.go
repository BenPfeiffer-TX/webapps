package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var templates = template.Must(template.ParseFiles("static/index.html", "static/create.html", "static/delete.html"))

type StatusMap struct {
	Name   string
	Status string
}

func getStatus() []StatusMap {
	file, err := os.Open("status.json")
	var statusArray []StatusMap
	if err != nil {
		fmt.Println("Error opening the json!" + err.Error())
		return statusArray
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&statusArray)
	if err != nil {
		fmt.Println("Error decoding the json!" + err.Error())
		return statusArray
	}

	return statusArray
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	statusArr := getStatus()
	data := map[string][]StatusMap{"StatusMap": statusArr}
	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	//display create.html, which contains a text box for a name and a 'submit' button
	statusArr := getStatus()
	data := map[string][]StatusMap{"StatusMap": statusArr}
	err := templates.ExecuteTemplate(w, "create.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	//can implement name restrictions here, or in create.html template
	newEntry := StatusMap{Name: name, Status: "Available"}

	//read the contents of status.json
	statusArray := getStatus()
	statusArray = append(statusArray, newEntry)
	file, err := os.OpenFile("status.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	//encode new data to status.json
	err = json.NewEncoder(file).Encode(statusArray)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	//display delete.html, which displays boxes of usernames and, when clicked, they are removed from list
	statusArr := getStatus()
	data := map[string][]StatusMap{"StatusMap": statusArr}
	err := templates.ExecuteTemplate(w, "delete.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func removeHandler(w http.ResponseWriter, r *http.Request) {
	//name := r.FormValue("name")
	//statusArray := getStatus()
	//...

	http.Redirect(w, r, "/", http.StatusFound)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {

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
