package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", home).Methods("GET")
	fmt.Println("server run at port 8080")
	http.ListenAndServe(":8080", r)

}
func home(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set
	fmt.Fprintf(w, r.Header.Get("name"))
}
