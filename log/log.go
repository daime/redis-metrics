package log

import (
	"log"
)

func Printf(fmt string, args ...interface{}) {
	log.Printf(fmt, args)
}
