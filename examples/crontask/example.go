package main

import (
	"log"
	"time"

	"github.com/aleasoluciones/goaleasoluciones/crontask"
)

func main() {
	crontask.New(func() { log.Println("CronTasks") }, "* * * * * * *").Run()
	time.Sleep(2 * time.Minute)
}
