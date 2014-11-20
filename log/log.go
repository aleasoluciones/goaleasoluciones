package log

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		for s := range c {
			switch {
			case s == syscall.SIGUSR1:
				EnableLogging()
				log.Println("Logging enabled")
			case s == syscall.SIGUSR2:
				log.Println("Logging disabled")
				DisableLogging()
			}
		}
	}()
}

func DisableLogging() {
	log.SetOutput(ioutil.Discard)
}

func EnableLogging() {
	log.SetOutput(os.Stdout)
}
