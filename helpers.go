package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Check to see if program index is valid.
func getProgramIndex(indexString string) (int, int, error) {
	// Try to convert it from a string to an integer.
	programIndex, err := strconv.Atoi(indexString)

	// Was it a valid index?
	if err == nil {
		// Make sure the program index is valid.
		var numPrograms = len(config.Programs)
		if programIndex >= 0 && programIndex < numPrograms {
			// Program index is valid.
			return programIndex, http.StatusOK, nil
		} else {
			// Program index is out of bounds.
			var errMsg = "Program index is out of bounds: " + indexString
			return programIndex, http.StatusNotFound,
				errors.New(errMsg)
		}
	} else {
		// It was an invalid program index.
		return -1, http.StatusBadRequest,
			errors.New("Invalid program index: " + indexString)
	}
}

// Check to see if step index is valid.
func getStepIndex(programIndexString string,
	stepIndexString string) (int, int, error) {
	// Get the program index.
	programIndex, statusCode, programErr := getProgramIndex(programIndexString)

	// Was it a valid program index?
	if programErr == nil {
		// Try to convert the step index from a string to an integer.
		stepIndex, stepErr := strconv.Atoi(stepIndexString)

		// Was it a valid step index?
		if stepErr == nil {
			// Make sure the step index is within the bounds of our array.
			var numSteps = len(config.Programs[programIndex].Steps)
			if stepIndex >= 0 && stepIndex < numSteps {
				// Step index is valid.
				return stepIndex, http.StatusOK, nil
			} else {
				// Step index is out of bounds.
				var errMsg = "Step index is out of bounds: " + stepIndexString
				return stepIndex, http.StatusNotFound,
					errors.New(errMsg)
			}
		} else {
			// It was an invalid step index.
			return -1, http.StatusBadRequest,
				errors.New("Invalid step index: " + stepIndexString)
		}
	} else {
		// It was an invalid program index.
		return programIndex, statusCode, programErr
	}
}

// Check to see if override index is valid.
func getOverrideIndex(indexString string) (int, int, error) {
	// Try to convert it from a string to an integer.
	overrideIndex, err := strconv.Atoi(indexString)

	// Was it a valid index?
	if err == nil {
		// Make sure the override index is valid.
		var numOverrides = len(config.Overrides)
		if overrideIndex >= 0 && overrideIndex < numOverrides {
			// Override index is valid.
			return overrideIndex, http.StatusOK, nil
		} else {
			// Override index is out of bounds.
			var errMsg = "Override index is out of bounds: " + indexString
			return overrideIndex, http.StatusNotFound,
				errors.New(errMsg)
		}
	} else {
		// It was an invalid override index.
		return -1, http.StatusBadRequest,
			errors.New("Invalid override index: " + indexString)
	}
}

// Check to see if JSON body is valid.
func parseBody(r *http.Request, dest interface{}) (int, error) {
	// Read the body of the request.
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	// Were we able to read the request?
	if err == nil {
		// We were able to read it, so let's see if the data was valid.
		err = json.Unmarshal(body, &dest)

		// Was the data valid?
		if err == nil {
			// Data was valid.
			return http.StatusOK, nil
		} else {
			// Data was invalid, so let the user know.
			return http.StatusBadRequest, errors.New("Malformed request body.")
		}
	} else {
		// There was an error. Send our error as the response.
		return http.StatusBadRequest,
			errors.New("Could not read request. Size of request might be too large.")
	}
}

// Update deployment counter and configuration file.
func updateDeployment() {
	// First, we need to increment the deployment counter.
	config.DeploymentCounter++

	// After we've incremented the deployment counter, we need to save the
	// actual deployment to the file.
	configFile, err := os.Create(configFilePath)

	// Defer closing the file until we're done with the function.
	defer configFile.Close()

	// Were we able to open the file?
	if err == nil {
		// We were able to open the file, so let's write to it.
		json.NewEncoder(configFile).Encode(config)
	} else {
		// There was an error.
		log.Printf(
			"ERROR: Could not save deployment to file: \"%s\"\n\tFile error: %v\n",
			configFilePath,
			err,
		)
	}
}
