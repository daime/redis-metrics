package main

import (
	"log"
	"time"

	"github.com/daime/redis-metrics/configuration"
	"github.com/daime/redis-metrics/redis"
	"github.com/daime/redis-metrics/signal"
	"github.com/daime/redis-metrics/statsd"
	"github.com/daime/redis-metrics/util"
)

func main() {
	signal.Handle()

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
	metricsSet := util.NewSet()
	metricsSet.AppendAll(config.Metrics)

	// Create a map to store only matching metrics
	replies := make(map[string]float64, len(config.Metrics))

	infoRequest := redis.NewInfoRequest(address)
	infoResponse, err := infoRequest.Send()
	if err != nil {
		log.Printf("Error getting redis info: %v", err)
		return
	}

	for metric, value := range infoResponse.Metrics {
		if metricsSet.Contains(metric) {
			replies[metric] = value
		}
	}

	s := statsd.NewStatsd(config.Statsd.Host, config.Statsd.Port)
	s.BulkGauge(address, replies)
}
