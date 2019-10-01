package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	router, err := rest.MakeRouter(
		rest.Get("/hello", goRestHandler),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	fmt.Println("server start at 8080")
	http.ListenAndServe(":"+strconv.Itoa(8080), api.MakeHandler())
}
func goRestHandler(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(map[string]string{"Body": req.Header.Get("name")})
}
