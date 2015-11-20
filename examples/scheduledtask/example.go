package main

import (
	"log"
	"time"

	"github.com/aleasoluciones/goaleasoluciones/scheduledtask"
)

func main() {

	t1 := scheduledtask.NewScheduledTask(func() { log.Println("taskWithOutTTL") }, 1*time.Second, 0)
	time.Sleep(5 * time.Second)
	log.Println("Stoping task")
	t1.Stop()

	// Example of task with ttl
	scheduledtask.NewScheduledTask(func() { log.Println("taskWithTTL") }, 1*time.Second, 5*time.Second)

	time.Sleep(10 * time.Second)
}
