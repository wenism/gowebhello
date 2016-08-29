package main

import (
	"net/http"
	"fmt"
	"os"
)

// This is injected at build time
var version = "undefined"

// This is injected at runtime
var environment = os.Getenv("ENVIRONMENT") 	
var colour = os.Getenv("COLOUR") 
var cport = os.Getenv("CPORT") 

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, 
		"<h1>Fancy Hello World Web App in Go</h1>" +
		"<h2>Environment: <span style='background-color: %s'>%s</span></h2>" +
		"<h2>Version: <span>%s</span></h2>",
		colour, environment, version)
}

func main() {
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/", handleIndex)

	http.ListenAndServe(":"+cport, nil)
}