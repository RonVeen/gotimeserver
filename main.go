package main

import (
	// Import the gorilla/mux library we just installed
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Declare a new router
	r := mux.NewRouter()

	// This is where the router is useful, it allows us to declare methods that
	// this path will be valid for
	r.HandleFunc("/t", timeHandler).Methods("GET")
	r.HandleFunc("/d", dateHandler).Methods("GET")
	r.HandleFunc("/f", datetimeHandler).Methods("GET")

	// We can then pass our router (after declaring all our routes) to this method
	// (where previously, we were leaving the second argument as nil)
	http.ListenAndServe(":8080", r)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	fmt.Fprintf(w, dt.Format("15:04:05"))
}

func dateHandler(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	fmt.Fprintf(w, dt.Format("02-01-2006"))
}

func datetimeHandler(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	fmt.Fprintf(w, dt.Format("02-01-2006 15:04:05"))
}