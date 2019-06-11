package log

import (
	"github.com/getsentry/raven-go"
	"log"
	"os"
)

var activeSentry bool = true

func InitializeSentry() {
	sentry := os.Getenv("SENTRY_DSN")
	if sentry == "" {
		activeSentry = false
		log.Println("===> Error: Sentry DSN environment not provisoned")
	}
	raven.SetDSN(sentry)
}

func LogError2Sentry(err error) {
	if activeSentry {
		raven.CaptureError(err, nil)
	} else {
		log.Println("===> Error: Sentry DSN environment not provisoned. Error received:", err)
	}
}
