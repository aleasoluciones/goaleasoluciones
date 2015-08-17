// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package periodictask

import (
	"log"
	"time"

	"github.com/gorhill/cronexpr"
)

type PeriodicTask struct {
	task     func()
	cronTime string
}

func New(task func(), cronTime string) *PeriodicTask {
	pt := PeriodicTask{task, cronTime}
	return &pt
}

func (pt *PeriodicTask) Run() {
	go func() {
		for {
			nextTime := cronexpr.MustParse(pt.cronTime).Next(time.Now())
			log.Println("Next execution", nextTime, pt.task)
			time.Sleep(nextTime.Sub(time.Now()))
			log.Println("Execution start")
			pt.task()
			log.Println("Execution end")
		}
	}()
}
