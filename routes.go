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
		"ProgramsListShow",
		"GET",
		"/programs",
		ProgramsListShow,
	},
	Route{
		"ProgramShow",
		"GET",
		"/programs/{programIndex}",
		ProgramShow,
	},
	Route{
		"StepsListShow",
		"GET",
		"/programs/{programIndex}/steps",
		StepsListShow,
	},
	Route{
		"StepShow",
		"GET",
		"/programs/{programIndex}/steps/{stepIndex}",
		StepShow,
	},
	Route{
		"OverridesListShow",
		"GET",
		"/overrides",
		OverridesListShow,
	},
	Route{
		"OverrideShow",
		"GET",
		"/overrides/{overrideIndex}",
		OverrideShow,
	},
	Route{
		"ConfigUpdate",
		"POST",
		"/config",
		ConfigUpdate,
	},
	Route{
		"ProgramsListUpdate",
		"POST",
		"/programs",
		ProgramsListUpdate,
	},
	Route{
		"ProgramUpdate",
		"POST",
		"/programs/{programIndex}",
		ProgramUpdate,
	},
	Route{
		"StepsListUpdate",
		"POST",
		"/programs/{programIndex}/steps",
		StepsListUpdate,
	},
	Route{
		"StepUpdate",
		"POST",
		"/programs/{programIndex}/steps/{stepIndex}",
		StepUpdate,
	},
	Route{
		"OverridesListUpdate",
		"POST",
		"/overrides",
		OverridesListUpdate,
	},
	Route{
		"OverrideUpdate",
		"POST",
		"/overrides/{overrideIndex}",
		OverrideUpdate,
	},
}
