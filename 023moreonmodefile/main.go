package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Go mod from Go")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome) // Note pass the reference of the function i.e. function Name

	log.Fatal(http.ListenAndServe(":8000", r))
}

func greeter() {
	fmt.Println("Hey mod there is greeter")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Go</h1>"))
}
