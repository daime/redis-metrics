package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/daime/redis-metrics/redis"
)

type configuration struct {
	Interval  int64    `json:"interval"`
	Addresses []string `json:"addresses"`
	Metrics   []string `json:"metrics"`
}

func main() {
	config := readConfiguration("config.json")

	tickerChannel := time.NewTicker(time.Second * time.Duration(config.Interval)).C

	for {
		select {
		case <-tickerChannel:
			for _, address := range config.Addresses {
				go info(address, config.Metrics)
			}
		}
	}
}

func readConfiguration(fileName string) configuration {
	var config configuration

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error loading configuration file from %s: %s", fileName, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error decoding JSON from %s: %s", fileName, err)
	}

	return config
}

func info(address string, metrics []string) {
	// Transform selected metrics slice to map
	metricsMap := make(map[string]bool, len(metrics))
	for _, metric := range metrics {
		metricsMap[metric] = true
	}

	// Create a map to store only matching metrics
	replies := make(map[string]float64, len(metrics))

	infoRequest := &redis.InfoRequest{
		Address: address,
	}
	infoResponse, err := infoRequest.Send()
	if err != nil {
		log.Printf("Error getting redis info: %v", err)
		return
	}

	for metric, value := range infoResponse.Metrics {
		if _, hasKey := metricsMap[metric]; hasKey {
			replies[metric] = value
		}
	}

	log.Printf("%s | %s => %v\n", address, "parsed at", infoResponse.ParsedAt)
	for metric, value := range replies {
		log.Printf("%s | %s => %f\n", address, metric, value)
	}
}
