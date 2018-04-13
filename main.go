package main

import (
	"net/http"
	"log"
	"os"
	"strings"
)

var Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	printLogo(Info)
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

func printLogo(logger *log.Logger) {
	for idx, l := range strings.Split(logo, "\n") {
		if idx < 4 {
			logger.Println("\x1b[33m" + l + "\x1b[39m")
		} else if idx < 6 {
			logger.Println("\x1b[32m" + l + "\x1b[39m")
		} else {
			logger.Println("\x1b[37m" + l + "\x1b[39m")
		}
	}
}

var logo = `
 _______  _______  __   __  _______  ___      ___      _______ 
|       ||       ||  | |  ||       ||   |    |   |    |       |
|    ___||   _   ||  |_|  ||    ___||   |    |   |    |   _   |
|   | __ |  | |  ||       ||   |___ |   |    |   |    |  | |  |
|   ||  ||  |_|  ||       ||    ___||   |___ |   |___ |  |_|  |
|   |_| ||       ||   _   ||   |___ |       ||       ||       |
|_______||_______||__| |__||_______||_______||_______||_______|
`