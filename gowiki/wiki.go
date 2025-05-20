package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

// each page in our wiki will have a title and a body
type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("template/edit.html", "template/view.html", "template/home.html"))
var validPath = regexp.MustCompile("^/(edit|save|view|delete|home|static)/([a-zA-Z0-9_-]+)$")

// this function validates the web path when accessing a page
// this is made obsolete by our handler function, makeHandler
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil
}

// this function wraps the handler calls, performing a validation as getTitle()
// but without requiring duplicate code in each handler func
func makeHandler(fn func(*log.Logger, http.ResponseWriter, *http.Request, string), l *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Here we will extract the page title from the Request and call the provided handler 'fn'
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(l, w, r, m[2])
	}
}

// this method addresses persistent storage of our Pages
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile("data/"+filename, p.Body, 0600)
}

// this method addresses deleting existing files
func (p *Page) dele() error {
	filename := p.Title + ".txt"
	return os.Remove("data/" + filename)
}

// this method is for loading a page from a file
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile("data/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, err
}

// this function executes arbitrary templates to our responsewriter
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// this handler allows us to view a wiki page, will handle URLs prefixed with "/view/"
func viewHandler(l *log.Logger, w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	l.Println(remoteIP(r) + " viewed " + title)
	renderTemplate(w, "view", p)
}

// this handler allows editing existing or new wiki pages
func editHandler(l *log.Logger, w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	l.Println(remoteIP(r) + " edited " + title)
	renderTemplate(w, "edit", p)
}

// this handler is for saving edits and new pages to the filesystem
func saveHandler(l *log.Logger, w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}

	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	l.Println(remoteIP(r) + " created " + title)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// this handler is for deleting existing pages
func deleHandler(l *log.Logger, w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		//page already doesnt exist
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	err = p.dele()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	l.Println(remoteIP(r) + " deleted " + title)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

// this handler is responsible for displaying the front page
func homeHandler(l *log.Logger, w http.ResponseWriter, r *http.Request, _ string) {
	l.Println(remoteIP(r) + " was redirected to home")
	http.Redirect(w, r, "/", http.StatusFound)
}

// this function is for determining the users IP address
func remoteIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		return xForwardedFor
	}
	ip := r.RemoteAddr
	return ip
}

func main() {
	//initialize log file and logger to write to
	file, err := os.OpenFile("logs/wiki.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: ", err)
	}
	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("----------")
	logger.Println("starting web server...")
	defer file.Close()

	staticDir := http.Dir("./static")
	fs := http.FileServer(staticDir)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/view/", makeHandler(viewHandler, logger))
	http.HandleFunc("/edit/", makeHandler(editHandler, logger))
	http.HandleFunc("/save/", makeHandler(saveHandler, logger))
	http.HandleFunc("/delete/", makeHandler(deleHandler, logger))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		renderTemplate(w, "home", nil)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
