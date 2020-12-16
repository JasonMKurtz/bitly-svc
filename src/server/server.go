package server

/*
	This is a simple wrapper around Golang's built-in HTTP server, which provides a simple solution for endpoint control as well as setting the port that the server will listen on.
	These configuration options are set in the main() method of endpoint.go
*/

import (
	"fmt"
	"net/http"
	"regexp"
)

type RouteHandler func(w http.ResponseWriter, r *http.Request, route Route)

type Route struct {
	Route   string
	Handler RouteHandler
}

type Routes struct {
	Routes []Route
	Port   string
}

func (r *Routes) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.Routes {
		reg, _ := regexp.Compile(route.Route)
		path := req.URL.Path
		if reg.MatchString(path) {
			w.Header().Set("Content-Type", "application/json")
			route.Handler(w, req, route)
			return
		}
	}

	http.NotFound(w, req)
	return
}

func (r *Routes) Listen() {
	fmt.Printf("Listening on :%s...\n", r.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", r.Port), r)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
