// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package periodictask

import (
	"log"
	"time"

	"github.com/gorhill/cronexpr"
)

func PeriodicTask(task func(), cronTime string) {
	for {
		nextTime := cronexpr.MustParse(cronTime).Next(time.Now())
		log.Println("Next execution", nextTime, task)
		time.Sleep(nextTime.Sub(time.Now()))
		log.Println("Execution start")
		task()
		log.Println("Execution end")
	}
}
