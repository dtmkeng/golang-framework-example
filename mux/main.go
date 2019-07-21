package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type student struct {
	Name string
	Age  int
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/list", showList)
	fmt.Println("server start at 8000 ...")
	log.Fatal(http.ListenAndServe(":8000", r))
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
