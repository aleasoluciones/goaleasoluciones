package scheduled_task

import (
	"github.com/aleasoluciones/goaleasoluciones/clock"
	"log"
	"time"
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

func sleepAndPrint(duration time.Duration, message string) {
	log.Println("sleeping ", duration)
	log.Println(message)
	time.Sleep(duration)
	log.Println("awake")
}

func main() {
	task := NewScheduledTask(
		func() {
			sleepAndPrint(3*time.Second, "Task 1")
		},
		5*time.Second,  // periode
		10*time.Second, // ttl
	)
	log.Println("Executing")
	time.Sleep(20 * time.Second)

	task2 := NewScheduledTask(
		func() {
			sleepAndPrint(3*time.Second, "Task 2")
		},
		1*time.Second,  // periode
		10*time.Second, // ttl
	)
	time.Sleep(10 * time.Second)
	task.Stop()
	task2.Stop()
	time.Sleep(20 * time.Second)
}
