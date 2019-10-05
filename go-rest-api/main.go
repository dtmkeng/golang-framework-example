package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
)

// NewRelicMiddleware ...
type NewRelicMiddleware struct{}

// MiddlewareFunc ...
func (m NewRelicMiddleware) MiddlewareFunc(next rest.HandlerFunc) rest.HandlerFunc {
	return func(writer rest.ResponseWriter, request *rest.Request) {
		time.Sleep(time.Millisecond * 200)
		next(writer, request)
		time.Sleep(time.Millisecond * 200)
	}
}
func main() {
	api := rest.NewApi()
	api.Use(&NewRelicMiddleware{})
	router, err := rest.MakeRouter(
		rest.Get("/", goRestHandler),
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
