package redis

import (
	r "github.com/garyburd/redigo/redis"
)

// Redis is the interface that holds the methods that can execute against redis
type Redis interface {
	Info() ([]byte, error)
}

type redis struct {
	address string
}

// NewRedis creates a struct filled with a redis instance address (host:port)
func NewRedis(address string) Redis {
	return &redis{
		address: address,
	}
}

// Info dials redis, executes INFO command and parses bulk string reply into a
// byte slice
// It's self contained about opening and closing a tcp connection to redis
func (redis *redis) Info() ([]byte, error) {
	connection, err := r.Dial("tcp", redis.address)
	if err != nil {
		return nil, err
	}
	defer connection.Close()

	reply, err := connection.Do("INFO")
	if err != nil {
		return nil, err
	}

	return reply.([]byte), nil
}
