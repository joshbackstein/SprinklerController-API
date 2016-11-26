package main

import (
	"errors"
	"net/http"
	"strconv"
)

// Check to see if program with ID exists.
func getProgramIndexByIdString(idString string) (int, int, error) {
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
			return programIndex, http.StatusOK, nil
		} else {
			// ID could not be found.
			return -1, http.StatusNotFound,
				errors.New("Could not find program with ID: " + idString)
		}
	} else {
		// It was an invalid ID.
		return -1, http.StatusBadRequest,
			errors.New("Invalid program ID: " + idString)
	}
}

// Check to see if step index is valid.
func getStepIndex(indexString string, programIndex int) (int, int, error) {
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
				return stepIndex, http.StatusOK, nil
			} else {
				// Step index is out of bounds.
				var errMsg = "Step index is out of bounds: " + indexString
				return stepIndex, http.StatusNotFound,
					errors.New(errMsg)
			}
		} else {
			// It was an invalid program index.
			return -1, http.StatusBadRequest,
				errors.New("Invalid program index: " + strconv.Itoa(programIndex))
		}
	} else {
		// It was an invalid ID.
		return -1, http.StatusBadRequest,
			errors.New("Invalid step index: " + indexString)
	}
}
