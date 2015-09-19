package redis

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	info := `
		# Memory
		used_memory:1000000
		used_memory_human:1.0M
		shouldNotBeParsed

		# Keyspace
		db0:keys=123,expires=37,avg_ttl=56789
		db1:keys=456,expires=48,avg_ttl=67890
	`

	lines := Parse([]byte(info))

	expected := map[string]float64{
		"used_memory": 1000000,
		"db0.keys":    123,
		"db0.expires": 37,
		"db0.avg_ttl": 56789,
		"db1.keys":    456,
		"db1.expires": 48,
		"db1.avg_ttl": 67890,
	}

	if !reflect.DeepEqual(expected, lines) {
		t.Errorf("Expected %v, but got %v", expected, lines)
	}
}
