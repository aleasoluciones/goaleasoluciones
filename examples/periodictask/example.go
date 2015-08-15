package main

import (
	"log"
	"time"

	"github.com/aleasoluciones/goaleasoluciones/periodictask"
)

func main() {
	periodictask.New(func() { log.Println("PeriodickTask") }, "* * * * * *")
	time.Sleep(2 * time.Minute)
}
