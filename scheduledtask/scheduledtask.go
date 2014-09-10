package scheduledtask

import (
	"log"
	"time"

	"github.com/aleasoluciones/goaleasoluciones/clock"
)

type ScheduledTask struct {
	periode time.Duration
	ttl     time.Duration
	task    func()
	finish  chan struct{}
	done    chan struct{}
	clock   clock.Clock
}

func NewScheduledTask(task func(), periode, ttl time.Duration) *ScheduledTask {
	scheduledTask := ScheduledTask{
		task:    task,
		periode: periode,
		ttl:     ttl,
		finish:  make(chan struct{}),
		done:    make(chan struct{}),
		clock:   clock.NewClock(),
	}
	go scheduledTask.run()
	return &scheduledTask
}

func (scheduler *ScheduledTask) run() {
	defer close(scheduler.done)
	scheduledUntil := scheduler.clock.Now().Add(scheduler.ttl)
	log.Println("Scheduling until", scheduledUntil)
	for {
		now := scheduler.clock.Now()
		if scheduler.ttl != 0 && now.After(scheduledUntil) {
			break
		}
		nextExecution := now.Add(scheduler.periode)

		scheduler.task()

		log.Println("Next execution ", nextExecution)
		log.Println("Waiting", (nextExecution.Sub(scheduler.clock.Now())))

		select {
		case <-time.After(nextExecution.Sub(scheduler.clock.Now())):
			continue
		case <-scheduler.finish:
			log.Println("Quiting")
			return
		}
	}
	log.Println("TTL reached")
}

func (scheduler *ScheduledTask) Stop() {
	close(scheduler.finish)
	<-scheduler.done
}
