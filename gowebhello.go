package main

import (
	"net/http"
	"fmt"
	"os"
    "log"
    "log/syslog"
)

// This is injected at build time
var version = "undefined"
var goVersion = "undefined"
var buildTime = "undefined"

// This is injected at runtime
var environment = os.Getenv("ENVIRONMENT") 	
var colour = os.Getenv("COLOUR") 
var cport = os.Getenv("CPORT") 

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy")
    log.Print("responding to /health")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, 
		"<h1>Fancy Hello World Web App in Go</h1>" +
        "<br/><br/>"
        "<ul><li>Feature 1</li><li>Feature 2</li></ul>"+
        "<br/><br/>"
		"<h3>Environment: <span style='background-color: %s'>%s</span></h3>" +
		"<h3>Version: <span>%s</span></h3>" +
		"<h3>Built using: <span>%s</span></h3>" +
        "<h3>Built Time: <span>%s</span></h3>",
		colour, environment, version, goVersion, buildTime)
    log.Print("responding to /")
}

func setupLog() {
    // Configure logger to write to the syslog. You could do this in init(), too.
    logwriter, e := syslog.New(syslog.LOG_NOTICE, "gowebhello")
    if e == nil {
        log.SetOutput(logwriter)
    }
}

func main() {
    setupLog()
    
    log.Printf("Starting web server with ENVIRONMENT=%s COLOUR=%s CPORT=%s VERSION=%s GOVERSION=%s", environment, colour, cport, versionm buildTime)
    
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/", handleIndex)

	http.ListenAndServe(":"+cport, nil)
}