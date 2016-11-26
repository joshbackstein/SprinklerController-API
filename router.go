package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Create new router.
func NewRouter() *mux.Router {
	// Create a new router with our routes.
	router := mux.NewRouter().StrictSlash(true)

	// Loop through our routes and add them to the router.
	for _, route := range routes {
		// Create handler to log requests.
		var handler http.Handler = Logger(route.HandlerFunc, route.Name)

		// Create route to this handler.
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// Return router.
	return router
}
