package configuration

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration stores parsed data from JSON configuration file.
type Configuration struct {
	// Interval stores how often to query redis about its information.
	Interval int64 `json:"interval"`
	// Addresses stores a list of redis instances to query.
	Addresses []string `json:"addresses"`
	// Metrics stores a list of redis infomations to send to statsd.
	Metrics []string `json:"metrics"`
}

// Load reads, parses and returns data from JSON configuration file parsed
// within Configuration struct. If any error occurs opening file or decoding
// file content it fails and interrupts program execution.
func Load(fileName string) Configuration {
	var config Configuration

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening configuration file from %s: %s", fileName, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error decoding JSON from %s: %s", fileName, err)
	}

	return config
}
