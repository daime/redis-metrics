package parse

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const (
	separator         = ":"
	keyspaceSeparator = ","
	equal             = "="
)

// Parse reads redis INFO reply and parse it into map[string]float64.
// Only metrics with valid float64 values are eligible to store into result map.
// Sections like "Keyspace" are handled different than others once their values
// have slightly different format.
func Parse(reply []byte) map[string]float64 {
	statistics := make(map[string]float64)

	reader := bytes.NewReader(reply)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, separator) {
			line = strings.TrimSpace(line)
			parts := strings.Split(line, separator)
			key := parts[0]
			value := parts[1]

			// Parse and store only valid float64 values
			if value, err := strconv.ParseFloat(value, 64); err == nil {
				statistics[key] = value
				continue
			}

			// Check if value is formatted as "keys=0,expires=0,avg_ttl=0"
			if strings.Contains(value, keyspaceSeparator) {
				// Split value into string slice ["keys=0", "expires=0", "avg_ttl=0"]
				for _, value := range strings.Split(value, keyspaceSeparator) {
					// Check if each slice value is formatted as "keys=0"
					if strings.Contains(value, equal) {
						// Split value into string slice ["keys", "0"]
						parts := strings.Split(value, equal)
						// Parse and store only valid float64 values
						if value, err := strconv.ParseFloat(parts[1], 64); err == nil {
							// Create key based on previous key and part of each
							// slice. Example:
							// Source: db0:keys=0,expires=0,avg_ttl=0
							// Become:	db0.keys
							// 			db0.expires
							//			db0.avg_ttl
							composedKey := fmt.Sprintf("%s.%s", key, parts[0])
							statistics[composedKey] = value
						}
					}
				}
			}
		}
	}

	return statistics
}
