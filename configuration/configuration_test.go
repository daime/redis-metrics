package configuration_test

import (
	"github.com/daime/redis-metrics/configuration"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	expected := &configuration.Configuration{
		Interval: 1,
		Redis: []configuration.Redis{
			configuration.Redis{
				Host:  "redis_1",
				Port:  6379,
				Alias: "master",
			},
			configuration.Redis{
				Host:  "redis_2",
				Port:  6379,
				Alias: "slave",
			},
		},
		Addresses: []string{
			"redis_1:6379",
			"redis_2:6379",
		},
		Metrics: []string{
			"used_memory",
			"mem_fragmentation_ratio",
			"db0.keys",
		},
		Statsd: configuration.Statsd{
			Host: "statsd",
			Port: 8125,
		},
	}

	config := configuration.Load("../config.json")

	if !reflect.DeepEqual(expected, &config) {
		t.Errorf("Expected:\n\t%v\nbut got:\n\t%v\n", expected, config)
	}
}
