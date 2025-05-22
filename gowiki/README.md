# Web app 1: gowiki tutorial 📚💻
This first webapp is from me following this step by step tutorial on [go.dev](https://go.dev/doc/articles/wiki/). The below description shamelessly provided by llama3.1:8b-instruct-q8_0

This project was created to help me learn the basics of web development in Go. It allows users to create and edit wiki entries, 
providing a basic CRUD (Create, Read, Update, Delete) interface for managing content 💼.

## Key Takeaways 📝
---------------

Through building this application, I gained hands-on experience with:

*   **HTTP ResponseWriters**: Writing responses back to the client using `http.ResponseWriter` 📨.
*   **HTTP Handlers**: Defining functions that handle HTTP requests and respond accordingly ⚙️.
*   **Function Literals**: Using anonymous functions to implement simple logic within handlers 🔩.
*   **Basic HTML Formatting**: Rendering simple HTML templates for displaying data 💻.

## Features ✨
-----

* Create new wiki entries 📝
* Edit existing wiki entries ✍️
* Delete unwanted wiki entries 🚮
* Redirect to a home page 🏠  


## Running the Application 🎉
------------------------

To run the application, simply execute:

```bash
git clone https://github.com/BenPfeiffer-TX/webapps
cd webapps/gowiki
go build . && ./wiki
```

Then, navigate to `http://localhost:8080` in your web browser to access the wiki editor 🔗.

**Note**: This is a basic implementation and not intended for production use. However, it serves as a valuable learning experience for anyone looking to explore Go's capabilities in web development 🤓.
