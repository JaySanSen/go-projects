package main

import (
	"fmt"
	"net/http"
)

func main() {
	// A simple print statement
	fmt.Println("Hello world")


	/*
		http.HandleFunc is used to register a function that will be called when a specific HTTP request comes in.
		here the / specifies to listen to every request.


		w and r are the variable names
		http.ResponseWriter and http.Request are the types
		* in *http.Request denotes the passing of a reference rather than a copy of the object.
	*/
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from my web service")
	})
}
