package signal

import (
	"os"
	"os/signal"
	"syscall"
)

func Handle(fn func()) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGHUP)
	go func() {
		for _ = range channel {
			fn()
		}
	}()
}
