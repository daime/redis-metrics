package redis

import "time"

// InfoRequest stores information to query redis INFO command.
type InfoRequest struct {
	// Redis interface
	redis Redis
}

// InfoResponse stores parsed data obtained from the response of INFO command.
type InfoResponse struct {
	// Metrics stores parsed response from INFO command.
	// Example: map[used_memory:1234.5678]
	Metrics map[string]float64
	// ParsedAt stores the UTC instant that INFO command was queried and parsed.
	ParsedAt time.Time
}

// NewInfoRequest receives a redis instance address (host:port) and returns a
// InfoRequest struct filled with Redis interface pointing to that address
func NewInfoRequest(address string) *InfoRequest {
	return &InfoRequest{
		redis: NewRedis(address),
	}
}

// Send executes redis INFO command, parses the response and returns it
// using InfoResponse struct. If any error occurs connecting or querying redis
// it returns nil InfoResponse and an error filled with the cause.
func (req *InfoRequest) Send() (res *InfoResponse, err error) {
	reply, err := req.redis.Info()
	if err != nil {
		return res, err
	}

	metrics := Parse(reply)

	res = &InfoResponse{
		Metrics:  metrics,
		ParsedAt: time.Now().UTC(),
	}
	return res, nil
}
