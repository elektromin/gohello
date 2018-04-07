package main

import (
	"net/http"
	"log"
	"os"
)

var Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/api/station/closest", closestStation)
	http.ListenAndServe(":8080", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	Info.Println("ping")
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}
