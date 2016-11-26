package main

import (
	"log"
	"net/http"
	"time"
)

// Create logger to log HTTP requests.
func Logger(inner http.Handler, name string) http.Handler {
	// Return an HTTP handler function that will log requests as they are made.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the current time.
		start := time.Now()

		// Serve the HTTP up.
		inner.ServeHTTP(w, r)

		// Log the request.
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
