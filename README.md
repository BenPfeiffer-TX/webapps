# Ben's web apps repository
In this public github repo I am storing my attempts at developing web applications in Go. This is intended to demonstrate my growth and understanding of the concepts behind web application development and practical usage of Go

## Web app 1: gowiki tutorial 📚💻
This first webapp is from me following this step by step tutorial on [go.dev](https://go.dev/doc/articles/wiki/). The below description shamelessly provided by llama3.1:8b-instruct-q8_0

This project was created to help me learn the basics of web development in Go. It allows users to create and edit wiki entries, 
providing a basic CRUD (Create, Read, Update, Delete) interface for managing content 💼.

### Key Takeaways 📝
---------------

Through building this application, I gained hands-on experience with:

*   **HTTP ResponseWriters**: Writing responses back to the client using `http.ResponseWriter` 📨.
*   **HTTP Handlers**: Defining functions that handle HTTP requests and respond accordingly ⚙️.
*   **Function Literals**: Using anonymous functions to implement simple logic within handlers 🔩.
*   **Basic HTML Formatting**: Rendering simple HTML templates for displaying data 💻.

### Features ✨
-----

*   Create new wiki entries 📝
*   Edit existing wiki entries ✍️
*   Delete unwanted wiki entries 🚮

### TODO list of features
-----
* implement delete functionality
* Store templates in tmpl/ and page data in data/.
* Add a handler to make the web root redirect to /view/FrontPage.
* Spruce up the page templates by making them valid HTML and adding some CSS rules.
* Implement inter-page linking by converting instances of [PageName] to <a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this)

### Running the Application 🎉
------------------------

To run the application, simply execute:

```bash
go build wiki.go && ./wiki
```

Then, navigate to `http://localhost:8080` in your web browser to access the wiki editor 🔗.

**Note**: This is a basic implementation and not intended for production use. However, it serves as a valuable learning experience for anyone looking to explore Go's capabilities in web development 🤓.
