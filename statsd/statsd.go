package statsd

import (
	"fmt"
	"net"
	"strings"
)

// Statsd interface
type Statsd interface {
	BulkGauge(redisAddress string, metrics map[string]float64) error
}

type statsd struct {
	host string
	port int
}

// NewStatsd func
func NewStatsd(host string, port int) Statsd {
	return &statsd{
		host: host,
		port: port,
	}
}

// TODO remove redisAddress. Metrics should have names formatted at this point
func (statsd *statsd) BulkGauge(redisAddress string, metrics map[string]float64) error {
	address := fmt.Sprintf("%s:%v", statsd.host, statsd.port)
	conn, err := net.Dial("udp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	a := strings.Replace(redisAddress, ":", ".", 1)
	for name, value := range metrics {
		m := fmt.Sprintf("%s.%s:%f|g", a, name, value)
		_, err := fmt.Fprintf(conn, m)
		// TODO handle error
		if err != nil {
			panic(err)
		}
	}
	return nil
}
