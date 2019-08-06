package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/hello", goRestHandler),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	http.ListenAndServe(":"+strconv.Itoa(8080), api.MakeHandler())
}
func goRestHandler(w rest.ResponseWriter, req *rest.Request) {

	w.WriteJson(map[string]string{"Body": "Hello World!"})
}
