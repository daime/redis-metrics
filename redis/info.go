package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// InfoRequest stores information to query redis INFO command.
type InfoRequest struct {
	// Address where the redis instance is running (host:port).
	// Example: http://localhost:6379
	Address string
}

// InfoResponse stores parsed data obtained from the response of INFO command.
type InfoResponse struct {
	// Metrics stores parsed response from INFO command.
	// Example: map[used_memory:1234.5678]
	Metrics map[string]float64
	// ParsedAt stores the UTC instant that INFO command was queried and parsed.
	ParsedAt time.Time
}

// Send executes redis INFO command, parses the response and returns it
// using InfoResponse struct. If any error occurs connecting or querying redis
// it returns nil InfoResponse and an error filled with the cause.
func (req *InfoRequest) Send() (res *InfoResponse, err error) {
	connection, err := redis.Dial("tcp", req.Address)
	if err != nil {
		return res, err
	}
	defer connection.Close()

	reply, err := connection.Do("INFO")
	if err != nil {
		return res, err
	}

	metrics := Parse(reply.([]byte))

	res = &InfoResponse{
		Metrics:  metrics,
		ParsedAt: time.Now().UTC(),
	}
	return res, nil
}
