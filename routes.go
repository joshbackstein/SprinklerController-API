package main

import (
	"net/http"
)

// Define route type.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

// Create routes.
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ConfigShow",
		"GET",
		"/config",
		ConfigShow,
	},
	Route{
		"ProgramsIndex",
		"GET",
		"/programs",
		ProgramsIndex,
	},
	Route{
		"ProgramShow",
		"GET",
		"/programs/{programId}",
		ProgramShow,
	},
	Route{
		"StepsIndex",
		"GET",
		"/programs/{programId}/steps",
		StepsIndex,
	},
	Route{
		"StepShow",
		"GET",
		"/programs/{programId}/steps/{stepIndex}",
		StepShow,
	},
}
