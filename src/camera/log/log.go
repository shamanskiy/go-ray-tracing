package log

import "log"

func Printf(formattedLog string, args ...any) {
	log.Printf(formattedLog, args...)
}
