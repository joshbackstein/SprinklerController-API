package main

// Types for config.
type Master struct {
	Enabled bool `json:"enabled"`
}
type Step struct {
	Zones     []bool `json:"zones"`
	StartTime int    `json:"startTime"`
	Duration  int    `json:"duration"`
}
type Program struct {
	Enabled    bool   `json:"enabled"`
	Name       string `json:"name"`
	DaysOfWeek []bool `json:"daysOfWeek"`
	Steps      []Step `json:"steps"`
}
type Override struct {
	Enabled        bool `json:"enabled"`
	UntilTurnedOff bool `json:"untilTurnedOff"`
	Duration       int  `json:"duration"`
}
type Config struct {
	DeploymentCounter int        `json:"deploymentCounter"`
	LocalHost         string     `json:"localHost"`
	LocalPort         int        `json:"localPort"`
	Host              string     `json:"host"`
	Port              int        `json:"port"`
	Master            Master     `json:"master"`
	Programs          []Program  `json:"programs"`
	Overrides         []Override `json:"overrides"`
}

// Types for JSON responses.
type Error struct {
	Error string `json:"error"`
}
type Ready struct {
	DeploymentCounter int `json:"deploymentCounter"`
}
