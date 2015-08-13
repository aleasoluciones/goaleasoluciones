// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

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

func NewScheduledTaskWithClock(task func(), periode, ttl time.Duration, clock clock.Clock) *ScheduledTask {
	scheduledTask := ScheduledTask{
		task:    task,
		periode: periode,
		ttl:     ttl,
		finish:  make(chan struct{}),
		done:    make(chan struct{}),
		clock:   clock,
	}
	go scheduledTask.run()
	return &scheduledTask
}
func NewScheduledTask(task func(), periode, ttl time.Duration) *ScheduledTask {
	return NewScheduledTaskWithClock(task, periode, ttl, clock.NewClock())
}

func (scheduler *ScheduledTask) run() {
	defer close(scheduler.done)
	scheduledUntil := scheduler.clock.Now().Add(scheduler.ttl)
	for {
		now := scheduler.clock.Now()
		if scheduler.ttl != 0 && now.After(scheduledUntil) {
			break
		}
		nextExecution := now.Add(scheduler.periode)

		scheduler.task()

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
