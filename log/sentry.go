package log

import (
	"github.com/getsentry/raven-go"
	"log"
	"os"
)

var isSentryActive bool = false

func InitializeSentry() {
	sentry := os.Getenv("SENTRY_DSN")
	if sentry == "" {
		isSentryActive = false
		log.Println("===> Error: Sentry DSN environment not provisoned")
	}
	raven.SetDSN(sentry)
}

func LogError2Sentry(err error) {
	if isSentryActive {
		raven.CaptureError(err, nil)
	} else {
		log.Println("===> Error: Sentry DSN environment not provisoned. Error received:", err)
	}
}
