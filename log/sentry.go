package log

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

// isSentryActive initialization was turned to false to disable sentry
var isSentryActive bool = false

func InitializeSentry() {
	sentryDsn := os.Getenv("SENTRY_DSN")
	if sentryDsn == "" {
		isSentryActive = false
		log.Println("===> Error: Sentry DSN environment not provisoned")
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDsn,
	})

	if err != nil {
		log.Printf("===> Error: Sentry initialization failed: %v\n", err)
	}
}

func InitSentry() {
	sentryDsn := os.Getenv("SENTRY_DSN")
	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDsn,
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

func LogError2Sentry(err error) {
	if isSentryActive {
		sentry.CaptureException(err)
	} else {
		log.Println("===> Error: Sentry DSN environment not provisoned. Error received:", err)
	}
}
