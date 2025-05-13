package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// each page in our wiki will have a title and a body
type Page struct {
	Title string
	Body  []byte
}

// this method addresses persistent storage of our Pages
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// this method is for loading a page from a file
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, err
}

// this handler allows us to view a wiki page, will handle URLs prefixed with "/view/"
func viewHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	//error handling for loadPage call
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	//p1 := &Page{Title: "TestPage1", Body: []byte("this is an example")}
	//p1.save()
	//p2, _ := loadPage("TestPage1")
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
