package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/daime/redis-metrics/redis"
)

func main() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range signalChannel {
			log.Printf("Catch you later: %s\n", sig)
			os.Exit(0)
		}
	}()

	config := readConfiguration("config.json")

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
