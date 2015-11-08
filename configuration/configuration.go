package configuration

import (
	"encoding/json"
	"log"
	"os"
)

type Redis struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Alias string `json:"alias"`
}

type Statsd struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// Configuration stores parsed data from JSON configuration file.
type Configuration struct {
	// Interval stores how often to query redis about its information.
	Interval int64 `json:"interval"`
	// Redis stores host port and the aliases for the redis instances.
	Redis []Redis `json:"redis"`
	// Addresses stores a list of redis instances to query.
	Addresses []string `json:"addresses"`
	// Metrics stores a list of redis infomations to send to statsd.
	Metrics []string `json:"metrics"`
	// Statsd stores host and port that will receive the metrics.
	Statsd Statsd `json:"statsd"`
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
