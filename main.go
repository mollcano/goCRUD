package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func items(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello items!")
}
func newItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
	}
	// handle the POST
	//Call to ParseForm makes form fields available.
	err := r.ParseForm()
	if err != nil {
		// Handle error here via logging and then return
		http.Error(w, "form parse failed", http.StatusInternalServerError)
	}
	name := r.PostFormValue("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	fmt.Println("hello World")
	http.HandleFunc("/", hello)
	http.HandleFunc("/items", items)
	http.HandleFunc("/new", newItem)
	http.ListenAndServe(":8000", nil)
}
