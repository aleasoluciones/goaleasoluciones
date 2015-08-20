package log

import (
	"github.com/getsentry/raven-go"
	"log"
	"os"
)

func InitializeSentry() {
	sentry := os.Getenv("SENTRY_DSN")
	if sentry == "" {
		log.Panic("Error: Sentry DSN environment not provisoned")
	}
	raven.SetDSN(sentry)
}

func LogError2Sentry(err error) {
	raven.CaptureError(err, nil)
}
