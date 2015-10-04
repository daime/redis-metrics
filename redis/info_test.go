package redis

import (
	"errors"
	"reflect"
	"testing"
)

const (
	address = "localhost:6379"
)

func TestSend(t *testing.T) {
	reply := "used_memory:1234.56"

	expectedMetrics := make(map[string]float64, 1)
	expectedMetrics["used_memory"] = 1234.56

	infoRequest := NewInfoRequest(address)
	infoRequest.redis = NewFakeRedis([]byte(reply), nil)

	infoResponse, err := infoRequest.Send()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	if !reflect.DeepEqual(expectedMetrics, infoResponse.Metrics) {
		t.Errorf("Expected %v but got %v", expectedMetrics, infoResponse.Metrics)
	}
}

func TestInfoReturnError(t *testing.T) {
	expectedError := errors.New("Fake redis error")

	infoRequest := NewInfoRequest(address)
	infoRequest.redis = NewFakeRedis(nil, expectedError)

	infoResponse, err := infoRequest.Send()
	if err == nil {
		t.Errorf("Expected nil but got %v", expectedError)
	}

	if infoResponse != nil {
		t.Errorf("Expected nil but got %v", infoResponse)
	}
}
