package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	mux := mux.NewRouter()
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/greets/{name}", greeting)

	fmt.Println("starting....")
	http.ListenAndServe(":"+port, mux)
}

func greeting(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	vars := mux.Vars(req)
	name := vars["name"]

	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{"message":"Hello %s"}`, name)
}

func homepage(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{"message":"welcome to the homepage!"}`)
}
