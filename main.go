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
	conf := readConfigurations("config.json")
	tickerChannel := time.NewTicker(time.Second * time.Duration(conf.Interval)).C

	for {
		select {
		case <-tickerChannel:
			go info(conf)
		}
	}
}

func ReadConfigurations(fileName string) configuration {
	var conf configuration
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	return conf
}

func info(c configuration) {
	// TODO make it multiple addresses
	address := c.Addresses[0]

	connection, err := redis.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Error dialing redis on address %s: %s", address, err)
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
