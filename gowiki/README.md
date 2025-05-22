# Web app 1: gowiki tutorial 📚💻
This first webapp is from me following this step by step tutorial on [go.dev](https://go.dev/doc/articles/wiki/). The below description shamelessly provided by llama3.1:8b-instruct-q8_0

This project was created to help me learn the basics of web development in Go. It allows users to create and edit wiki entries, 
providing a basic CRUD (Create, Read, Update, Delete) interface for managing content 💼.

## Key Takeaways 📝
---------------

Through building this application, I gained hands-on experience with:

*   **HTTP Servers**: Creating handlers for HTTP requests and managing HTML template rendering 🖥.
*   **Simple User Activity Logging**: Setting up a log file to track user activity based on IP 🕵️.
*   **Basic HTML Formatting**: Rendering simple HTML templates for displaying data 💻.

## Features ✨
-----

* Create new wiki entries 📝
* Edit existing wiki entries ✍️
* Delete unwanted wiki entries 🚮
* Redirect to a home page 🏠 
* Dynamic linking to entries
* CSS styling for web pages 


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
