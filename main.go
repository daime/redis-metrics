package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

var c = `{
	"interval": 1,
	"addresses": [
		"redis:6379"
	],
	"metrics": [
		"used_memory"
	]
}`

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
			go info(config)
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

func info(c configuration) {
	// TODO make it multiple addresses
	address := c.Addresses[0]

	connection, err := redis.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Error connecting to Redis: %s", err)
		return
	}

	reply, err := connection.Do("INFO")
	if err != nil {
		log.Fatalf("Error doing INFO command: %s", err)
		return
	}

	content := fmt.Sprintf("%s", reply)

	for _, line := range strings.Split(content, "\n") {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			fmt.Println(parts[0], parts[1])
		}
	}

	fmt.Println(c)
}
