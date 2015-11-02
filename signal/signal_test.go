package signal_test

import (
	"syscall"
	"testing"
	"time"

	"github.com/daime/redis-metrics/signal"
)

func receiveOnce(t *testing.T, done <-chan bool) {
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Error("Receive timeout")
	}
}
func Test_Handle_WithSIGHUP_ExecutesFn(t *testing.T) {
	done := make(chan bool, 1)
	signal.Handle(func() { done <- true })
	syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
	receiveOnce(t, done)
}

func Test_Handle_With2_SIGHUP_ExecutesFnTwice(t *testing.T) {
	done := make(chan bool, 1)
	value := 0
	signal.Handle(func() {
		done <- true
		value++
	})
	syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
	receiveOnce(t, done)
	syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
	receiveOnce(t, done)
	if value != 2 {
		t.Error("Handle not updated value")
	}
}
