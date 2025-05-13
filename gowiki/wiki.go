package main

import (
	"fmt"
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

func main() {
	p1 := &Page{Title: "TestPage1", Body: []byte("this is an example")}
	p1.save()
	p2, _ := loadPage("TestPage1")
	fmt.Println(string(p2.Body))
}
