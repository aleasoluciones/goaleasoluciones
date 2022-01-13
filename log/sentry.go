package log

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/raven-go"
	"github.com/getsentry/sentry-go"
)

// isSentryActive initialization was turned to false to disable sentry
var isSentryActive bool = false

// deprecated
func InitializeSentry() {
	sentry := os.Getenv("SENTRY_DSN")
	if sentry == "" {
		isSentryActive = false
		log.Println("===> Error: Sentry DSN environment not provisoned")
	}
	raven.SetDSN(sentry)
}

func InitSentry() {
	sentry_dsn := os.Getenv("SENTRY_DSN")
	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentry_dsn,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
		panic("Can't start sentry")
	}
}

func HandlePanic() {
	err := recover()
	if err != nil {
		sentry.CurrentHub().Recover(err)
		sentry.Flush(time.Second * 5)
	}
}

// deprecated
func LogError2Sentry(err error) {
	if isSentryActive {
		raven.CaptureError(err, nil)
	} else {
		log.Println("===> Error: Sentry DSN environment not provisoned. Error received:", err)
	}
}
