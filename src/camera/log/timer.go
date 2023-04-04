package log

import (
	"log"
	"time"
)

func TimeExecution(processName string) func() {
	start := time.Now()
	log.Printf("%s: start", processName)

	return func() {
		log.Printf("%s: end. Duration: %v", processName, time.Since(start))
	}
}
