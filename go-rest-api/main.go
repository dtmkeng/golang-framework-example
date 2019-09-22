package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
)

type NewRelicMiddleware struct{}

func (m NewRelicMiddleware) MiddlewareFunc(next rest.HandlerFunc) rest.HandlerFunc {
	return func(writer rest.ResponseWriter, request *rest.Request) {

		start := time.Now()
		next(writer, request)
		responseTime := time.Since(start)

		// Write it to the log
		fmt.Println(responseTime)

		// Make sure to pass the error back!
		// return err
	}
}
func main() {
	api := rest.NewApi()
	// api.Use(rest.DefaultDevStack...)
	// api.Use(&rest.AccessLogApacheMiddleware{})
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
