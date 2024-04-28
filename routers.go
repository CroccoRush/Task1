package main

import (
	"Task1/handlers"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"Ping",
		strings.ToUpper("Get"),
		"/ping",
		handlers.Ping,
	},

	Route{
		"Post",
		strings.ToUpper("Get"),
		"/set/{key}/{value}",
		handlers.Set,
	},

	Route{
		"Get",
		strings.ToUpper("Get"),
		"/get/{key}",
		handlers.Get,
	},
}
