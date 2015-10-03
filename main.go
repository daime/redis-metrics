package main

import (
	"log"
	"time"

	"github.com/daime/redis-metrics/configuration"
	"github.com/daime/redis-metrics/redis"
	"github.com/daime/redis-metrics/statsd"
)

func main() {
	config := configuration.Load("config.json")

	tickerChannel := time.NewTicker(time.Second * time.Duration(config.Interval)).C

	for {
		select {
		case <-tickerChannel:
			for _, address := range config.Addresses {
				go info(address, config)
			}
		}
	}
}

func info(address string, config configuration.Configuration) {
	// Transform selected metrics slice to map
	metricsMap := make(map[string]bool, len(config.Metrics))
	for _, metric := range config.Metrics {
		metricsMap[metric] = true
	}

	// Create a map to store only matching metrics
	replies := make(map[string]float64, len(config.Metrics))

	infoRequest := redis.NewInfoRequest(address)
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

	s := statsd.NewStatsd(config.Statsd.Host, config.Statsd.Port)
	s.BulkGauge(address, replies)
}
