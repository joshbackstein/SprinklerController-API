package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Index handler.
func Index(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Set status.
	w.WriteHeader(http.StatusOK)

	// Grab deployment counter to respond with.
	var ready Ready = Ready{
		DeploymentCounter: config.DeploymentCounter,
	}

	// Send our JSON response.
	json.NewEncoder(w).Encode(ready)
}

// ConfigShow handler.
func ConfigShow(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Set status.
	w.WriteHeader(http.StatusOK)

	// Send our JSON response.
	json.NewEncoder(w).Encode(config)
}

// MasterShow handler.
func MasterShow(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Set status.
	w.WriteHeader(http.StatusOK)

	// Send our JSON response.
	json.NewEncoder(w).Encode(config.Master)
}

// ProgramsListShow handler.
func ProgramsListShow(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Set status.
	w.WriteHeader(http.StatusOK)

	// Send our JSON response.
	json.NewEncoder(w).Encode(config.Programs)
}

// ProgramShow handler.
func ProgramShow(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get the program index.
	vars := mux.Vars(r)
	programIndexString := vars["programIndex"]
	programIndex, statusCode, err := getProgramIndex(programIndexString)

	// Set status.
	w.WriteHeader(statusCode)

	// Were we able to find it in our program array?
	if err == nil {
		// The program was found in the array.
		json.NewEncoder(w).Encode(config.Programs[programIndex])
	} else {
		// There was an error. Send it as the response.
		var e Error = Error{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
	}
}

// StepsListShow handler.
func StepsListShow(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get the program index.
	vars := mux.Vars(r)
	programIndexString := vars["programIndex"]
	programIndex, statusCode, err := getProgramIndex(programIndexString)

	// Set status.
	w.WriteHeader(statusCode)

	// Were we able to find it in our program array?
	if err == nil {
		// The program was found in the array.
		json.NewEncoder(w).Encode(config.Programs[programIndex].Steps)
	} else {
		// There was an error. Send it as the response.
		var e Error = Error{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
	}
}

// StepShow handler.
func StepShow(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get the program and step indices.
	vars := mux.Vars(r)
	programIndexString := vars["programIndex"]
	stepIndexString := vars["stepIndex"]
	programIndex, statusCode, programErr :=
		getProgramIndex(programIndexString)

	// Were we able to find the program in our program array?
	if programErr == nil {
		// The program was found in the array, but we still need to make sure the
		// step we're trying to access is valid.
		stepIndex, statusCode, stepErr :=
			getStepIndex(programIndexString, stepIndexString)

		// Set status.
		w.WriteHeader(statusCode)

		// Were we able to find the step in this program?
		if stepErr == nil {
			// We were able to find it within this program.
			json.NewEncoder(w).Encode(config.Programs[programIndex].Steps[stepIndex])
		} else {
			// There was an error. Send it as the response.
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

// OverridesListShow handler.
func OverridesListShow(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Set status.
	w.WriteHeader(http.StatusOK)

	// Send our JSON response.
	json.NewEncoder(w).Encode(config.Overrides)
}

// OverrideShow handler.
func OverrideShow(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get the override index.
	vars := mux.Vars(r)
	overrideIndexString := vars["overrideIndex"]
	overrideIndex, statusCode, err := getOverrideIndex(overrideIndexString)

	// Set status.
	w.WriteHeader(statusCode)

	// Were we able to find it in our override array?
	if err == nil {
		// The override was found in the array.
		json.NewEncoder(w).Encode(config.Overrides[overrideIndex])
	} else {
		// There was an error. Send it as the response.
		var e Error = Error{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(e)
	}
}
