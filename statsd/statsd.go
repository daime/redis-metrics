package statsd

import (
	"fmt"
	"net"
)

// Statsd interface
type Statsd interface {
	BulkGauge(metrics map[string]float64) error
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
func (statsd *statsd) BulkGauge(metrics map[string]float64) error {
	address := fmt.Sprintf("%s:%v", statsd.host, statsd.port)
	conn, err := net.Dial("udp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	for name, value := range metrics {
		m := fmt.Sprintf("%s:%f|g", name, value)
		_, err := fmt.Fprintf(conn, m)
		// TODO handle error
		if err != nil {
			panic(err)
		}
	}
	return nil
}
