package main

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
