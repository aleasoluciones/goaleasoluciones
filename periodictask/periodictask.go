// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package periodictask

import (
	"log"
	"time"

	"github.com/facebookgo/clock"
	"github.com/gorhill/cronexpr"
)

type PeriodicTask struct {
	task     func()
	cronTime string
	clock    clock.Clock
}

func NewWithClock(task func(), cronTime string, clock clock.Clock) *PeriodicTask {
	pt := PeriodicTask{task, cronTime, clock}
	go pt.run()
	return &pt
}

func New(task func(), cronTime string) *PeriodicTask {
	return NewWithClock(task, cronTime, clock.New())
}

func (pt *PeriodicTask) run() {
	for {
		nextTime := cronexpr.MustParse(pt.cronTime).Next(time.Now())
		log.Println("Next execution", nextTime, pt.task)
		time.Sleep(nextTime.Sub(time.Now()))
		log.Println("Execution start")
		pt.task()
		log.Println("Execution end")
	}
}
