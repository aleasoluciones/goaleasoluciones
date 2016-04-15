// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package crontask

import (
	"log"
	"time"

	"github.com/gorhill/cronexpr"
)

type CronTask struct {
	task     func()
	cronTime string
}

func New(task func(), cronTime string) *CronTask {
	return &CronTask{task, cronTime}
}

func (t *CronTask) Run() {
	go func() {
		for {
			nextTime := cronexpr.MustParse(t.cronTime).Next(time.Now())
			log.Println("Next execution", nextTime)
			time.Sleep(nextTime.Sub(time.Now()))
			log.Println("Execution start")
			t.task()
			log.Println("Execution end")
		}
	}()
}
