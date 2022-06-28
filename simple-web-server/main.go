package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is Not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello There!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() Error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST REQUEST SUCCESSFUL")
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	fmt.Fprintf(w, "First Name is %s and Last Name is %s \n", firstName, lastName)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080.\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
