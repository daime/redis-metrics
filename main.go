package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/daime/redis-metrics/configuration"
	"github.com/daime/redis-metrics/redis"
	"github.com/daime/redis-metrics/signal"
	"github.com/daime/redis-metrics/statsd"
	"github.com/daime/redis-metrics/util"
)

var (
	configFilePath string
)

func init() {
	flag.StringVar(&configFilePath, "config", "config.json", "Specify the configuration file path")
}

func main() {
	flag.Parse()

	log.Printf("Configuration loaded from file: %s", configFilePath)
	config := configuration.Load(configFilePath)
	signal.Handle(func() {
		config = configuration.Load("config.json")
		log.Printf("Reloading configuration from: %s", configFilePath)
	})

	tickerChannel := time.NewTicker(time.Second * time.Duration(config.Interval)).C

	for {
		select {
		case <-tickerChannel:
			for _, redisConfig := range config.Redis {
				address := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
				var alias = ""
				if redisConfig.Alias != "" {
					alias = redisConfig.Alias
				} else {
					alias = redisConfig.Host
				}
				go info(address, alias, config)
			}
		}
	}
}

func info(address string, alias string, config configuration.Configuration) {
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
			metricName := fmt.Sprintf("%s.%s", alias, metric)
			replies[metricName] = value
		}
	}

	s := statsd.NewStatsd(config.Statsd.Host, config.Statsd.Port)
	s.BulkGauge(replies)
}
