package configuration

import (
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	expected := &Configuration{
		Interval: 1,
		Addresses: []string{
			"redis_1:6379",
			"redis_2:6379",
		},
		Metrics: []string{
			"used_memory",
			"mem_fragmentation_ratio",
			"db0.keys",
		},
		Statsd: struct {
			Host string `json:"host"`
			Port int    `json:"port"`
		}{
			Host: "statsd",
			Port: 8125,
		},
	}

	config := Load("../config.json")

	if !reflect.DeepEqual(expected, &config) {
		t.Errorf("Expected:\n\t%v\nbut got:\n\t%v\n", expected, config)
	}
}
