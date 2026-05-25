package main

import (
	"fmt"      //to print and write in response writer
	"log"      //to log errors
	"net/http" //to create web server and routes
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "./static/forms.html")
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}
	//write the form data to the response
	fmt.Fprintf(w, "Form submitted\n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\nEmail: %s\nAddress: %s\n", name, email, address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not GET", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello World!")
}

// this is our main function
func main() {

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.Handle("/form", http.HandlerFunc(formHandler))
	http.Handle("/hello", http.HandlerFunc(helloHandler))

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
