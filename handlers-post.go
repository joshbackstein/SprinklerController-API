package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ConfigUpdate handler.
func ConfigUpdate(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Read request body.
	statusCode, err := parseBody(r, &config)

	// Set status code.
	w.WriteHeader(statusCode)

	// Was there an error?
	if err == nil {
		// It was parsed correctly, so increment the deployment counter and
		// save the new configuration file.
		updateDeployment()

		// Respond with JSON of new config.
		json.NewEncoder(w).Encode(config)
	} else {
		// It could not be parsed correctly.
		var e Error = Error{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
	}
}

// ProgramsListUpdate handler.
func ProgramsListUpdate(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Read request body.
	statusCode, err := parseBody(r, &config.Programs)

	// Set status code.
	w.WriteHeader(statusCode)

	// Was there an error?
	if err == nil {
		// It was parsed correctly, so increment the deployment counter and save
		// the new configuration file.
		updateDeployment()

		// Respond with JSON of new config.
		json.NewEncoder(w).Encode(config)
	} else {
		// It could not be parsed correctly.
		var e Error = Error{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
	}
}

// ProgramUpdate handler.
func ProgramUpdate(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get the program and step indices.
	vars := mux.Vars(r)
	programIndexString := vars["programIndex"]
	programIndex, statusCode, err := getProgramIndex(programIndexString)

	// Were we able to find the program in our program array?
	if err == nil {
		// The program was found in the array, so read the request body.
		statusCode, err := parseBody(r, &config.Programs[programIndex])

		// Set status code.
		w.WriteHeader(statusCode)

		// Was there an error?
		if err == nil {
			// It was parsed correctly, so increment the deployment counter and
			// save the new configuration file.
			updateDeployment()

			// Respond with JSON of new config.
			json.NewEncoder(w).Encode(config)
		} else {
			// It could not be parsed correctly.
			var e Error = Error{
				Error: err.Error(),
			}
			json.NewEncoder(w).Encode(e)
		}
	} else {
		// There was an error. Send it as the response.
		w.WriteHeader(statusCode)
		var e Error = Error{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
	}
}

// StepsListUpdate handler.
func StepsListUpdate(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get the program and step indices.
	vars := mux.Vars(r)
	programIndexString := vars["programIndex"]
	programIndex, statusCode, err := getProgramIndex(programIndexString)

	// Were we able to find the program in our program array?
	if err == nil {
		// The program was found in the array, so read the request body.
		statusCode, err := parseBody(r, &config.Programs[programIndex].Steps)

		// Set status code.
		w.WriteHeader(statusCode)

		// Was there an error?
		if err == nil {
			// It was parsed correctly, so increment the deployment counter and save
			// the new configuration file.
			updateDeployment()

			// Respond with JSON of new config.
			json.NewEncoder(w).Encode(config)
		} else {
			// It could not be parsed correctly.
			var e Error = Error{
				Error: err.Error(),
			}
			json.NewEncoder(w).Encode(e)
		}
	} else {
		// There was an error. Send it as the response.
		w.WriteHeader(statusCode)
		var e Error = Error{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
	}
}

// StepUpdate handler.
func StepUpdate(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get the program and step indices.
	vars := mux.Vars(r)
	programIndexString := vars["programIndex"]
	stepIndexString := vars["stepIndex"]
	programIndex, statusCode, programErr := getProgramIndex(programIndexString)

	// Were we able to find the program in our program array?
	if programErr == nil {
		// The program was found in the array, but we still need to make sure the
		// step we're trying to access is valid.
		stepIndex, statusCode, stepErr :=
			getStepIndex(programIndexString, stepIndexString)

		// Were we able to find the step in this program?
		if stepErr == nil {
			// We were able to find it within this program, so read the request body
			// to get the JSON.
			statusCode, stepErr :=
				parseBody(r, &config.Programs[programIndex].Steps[stepIndex])

			// Set status code.
			w.WriteHeader(statusCode)

			// Was it parsed correctly?
			if stepErr == nil {
				// It was parsed correctly, so increment the deployment counter and
				// save the new configuration file.
				updateDeployment()

				// Respond with JSON of new config.
				json.NewEncoder(w).Encode(config)
			} else {
				// It could not be parsed correctly.
				var e Error = Error{
					Error: stepErr.Error(),
				}
				json.NewEncoder(w).Encode(e)
			}
		} else {
			// There was an error. Send it as the response.
			w.WriteHeader(statusCode)
			var e Error = Error{
				Error: stepErr.Error(),
			}
			json.NewEncoder(w).Encode(e)
		}
	} else {
		// There was an error. Send it as the response.
		w.WriteHeader(statusCode)
		var e Error = Error{
			Error: programErr.Error(),
		}
		json.NewEncoder(w).Encode(e)
	}
}
