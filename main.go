package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// Global config variable.
var config Config
var configFilePath string = "./config.json"

// Main function.
func main() {
	// Load JSON file.
	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Configuration file opened.")

	// Load JSON file contents into config variable.
	json.Unmarshal(file, &config)
	fmt.Println("Configuration loaded.")

	// Blank line to make output look nicer.
	fmt.Println("")

	// Run HTTP server.
	runServer()
}

// Run the HTTP server.
func runServer() {
	// Figure out which port we want to listen on.
	var port string = strconv.Itoa(config.LocalPort)
	if port == "0" {
		// Port was invalid or not provided, so use port 4000
		// by default
		fmt.Println("WARNING: Port was invalid or not provided. Using port 4000.")
		port = "4000"
	}

	// We need some routes for our API.
	fmt.Println("Assigning HTTP route handlers.")
	router := NewRouter()

	// Start HTTP server.
	fmt.Println("Listening on " + config.LocalHost + ":" + port)
	http.ListenAndServe(config.LocalHost+":"+port, router)
}
