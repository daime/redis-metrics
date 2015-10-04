package signal

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Handle() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range signalChannel {
			log.Printf("Catch you later: %s\n", sig)
			os.Exit(0)
		}
	}()
}
