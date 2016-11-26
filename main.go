package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Types for config.
type Step struct {
	Zones     []bool `json:"zones"`
	StartTime int    `json:"startTime"`
	Duration  int    `json:"duration"`
}
type Program struct {
	Id         int    `json:"id"`
	Enabled    bool   `json:"enabled"`
	DaysOfWeek []bool `json:"daysOfWeek"`
	StepCount  int    `json:"stepCount"`
	Steps      []Step `json:"steps"`
}
type Config struct {
	DeploymentCounter int       `json:"deploymentCounter"`
	Url               string    `json:"url"`
	Port              int       `json:"port"`
	ProgramCount      int       `json:"programCount"`
	Programs          []Program `json:"programs"`
}

// Types for JSON responses.
type Error struct {
	Error string `json:"error"`
}
type Ready struct {
	DeploymentCounter int `json:"deploymentCounter"`
}

// Global config variable.
var config Config

// Main function.
func main() {
	// Load JSON file.
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Configuration file opened.")
	//fmt.Printf("%s\n", string(file))

	// Load JSON file contents into config variable.
	json.Unmarshal(file, &config)
	fmt.Println("Configuration loaded.")
	//fmt.Printf("Results: %v\n", config)

	// Blank line to make output look nicer.
	fmt.Println("")

	// Run HTTP server.
	runServer()
}

// Run the HTTP server.
func runServer() {
	// Figure out which port we want to listen on.
	var port string = strconv.Itoa(config.Port)
	if port == "0" {
		// Port was invalid or not provided, so use port 4000
		// by default
		fmt.Println("WARNING: Port was invalid or not provided. Using port 4000.")
		port = "4000"
	}

	// We need some routes for our API.
	fmt.Println("Assigning HTTP route handlers.")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/config", ConfigShow)
	router.HandleFunc("/programs", ProgramsIndex)
	router.HandleFunc("/programs/{programId}", ProgramShow)
	router.HandleFunc("/programs/{programId}/steps", StepsIndex)
	router.HandleFunc("/programs/{programId}/steps/{stepIndex}", StepShow)

	// Start HTTP server.
	fmt.Println("Listening on " + config.Url + ":" + port)
	http.ListenAndServe(config.Url+":"+port, router)
}

// Index handler.
func Index(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

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

	// Send our JSON response.
	json.NewEncoder(w).Encode(config)
}

// ProgramsIndex handler.
func ProgramsIndex(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Send our JSON response.
	json.NewEncoder(w).Encode(config.Programs)
}

// ProgramShow handler.
func ProgramShow(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get the program ID and the index in the array where that ID is located.
	vars := mux.Vars(r)
	programIdString := vars["programId"]
	programIndex, err := getProgramIndexByIdString(programIdString)

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

// StepsIndex handler.
func StepsIndex(w http.ResponseWriter, r *http.Request) {
	// Set return type.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get the program ID and the index in the array where that ID is located.
	vars := mux.Vars(r)
	programIdString := vars["programId"]
	programIndex, err := getProgramIndexByIdString(programIdString)

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

	// Get the program ID, the step ID, and the index in the array where those
	// IDs are located.
	vars := mux.Vars(r)
	programIdString := vars["programId"]
	stepIndexString := vars["stepIndex"]
	programIndex, programErr := getProgramIndexByIdString(programIdString)

	// Were we able to find the program in our program array?
	if programErr == nil {
		// The program was found in the array, but we still need to make sure the
		// step we're trying to access is valid.
		stepIndex, stepErr := getStepIndex(stepIndexString, programIndex)

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
		var e Error = Error{
			Error: programErr.Error(),
		}
		json.NewEncoder(w).Encode(e)
	}
}

// Check to see if program with ID exists.
func getProgramIndexByIdString(idString string) (int, error) {
	// Try to convert it from a string to an integer.
	programId, err := strconv.Atoi(idString)

	// Was it a valid ID?
	if err == nil {
		// Loop through programs until we find the one we're searching for.
		var programIndex int = -1
		for i := 0; i < len(config.Programs); i++ {
			if config.Programs[i].Id == programId {
				programIndex = i
			}
		}

		// If the program index is not -1, then we were able to find the program.
		if programIndex != -1 {
			// ID was found.
			return programIndex, nil
		} else {
			// ID could not be found.
			return -1, errors.New("Could not find program with ID: " + idString)
		}
	} else {
		// It was an invalid ID.
		return -1, errors.New("Invalid program ID: " + idString)
	}
}

// Check to see if step index is valid.
func getStepIndex(indexString string, programIndex int) (int, error) {
	// Try to convert it from a string to an integer.
	stepIndex, err := strconv.Atoi(indexString)

	// Was it a valid ID?
	if err == nil {
		// Make sure the program index is valid.
		if programIndex >= 0 {
			// Make sure the step index is within the bounds of our array.
			var numSteps = len(config.Programs[programIndex].Steps)
			if stepIndex >= 0 && stepIndex < numSteps {
				// Step index is valid.
				return stepIndex, nil
			} else {
				// Step index is out of bounds.
				var errMsg = "Step index is out of bounds: " + indexString
				return stepIndex, errors.New(errMsg)
			}
		} else {
			// It was an invalid program index.
			return -1, errors.New("Invalid program index: " + strconv.Itoa(programIndex))
		}
	} else {
		// It was an invalid ID.
		return -1, errors.New("Invalid step index: " + indexString)
	}
}
