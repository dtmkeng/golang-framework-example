package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type student struct {
	Name string
	Age  int
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		responseTime := time.Since(start)

		// Write it to the log
		fmt.Println(responseTime)

		// Make sure to pass the error back!
		// return err
		// next.ServeHTTP(w, r)
	})
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	// r.HandleFunc("/list", showList)
	r.Use(loggingMiddleware)

}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func showList(w http.ResponseWriter, r *http.Request) {
	s := &student{
		Name: "keng",
		Age:  15,
	}
	json.NewEncoder(w).Encode(s)
}
