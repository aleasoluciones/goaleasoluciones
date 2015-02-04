// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package retrier

import (
	"time"

	"github.com/aleasoluciones/goaleasoluciones/clock"
)

const (
	maximumAttempts = 5
	interval        = 30 * time.Second
)

type Retrier struct {
	sleeper         clock.Sleeper
	MaximumAttempts int
	Interval        time.Duration
}

func NewRetrierWithSleeper(sleeper clock.Sleeper) *Retrier {
	return &Retrier{
		sleeper:         sleeper,
		MaximumAttempts: maximumAttempts,
		Interval:        interval,
	}
}

func (r *Retrier) RunRetrying(wrapped func() (interface{}, error)) (interface{}, error) {
	attempts := 1
	result, err := wrapped()
	for err != nil && attempts < r.MaximumAttempts {
		attempts++
		r.sleeper.Sleep(r.Interval)
		result, err = wrapped()
	}
	return result, err
}
