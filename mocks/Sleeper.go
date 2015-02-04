// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package mocks

import "github.com/stretchr/testify/mock"

import "time"

type Sleeper struct {
	mock.Mock
}

func (m *Sleeper) Sleep(duration time.Duration) {
	m.Called(duration)
}
func (m *Sleeper) SleepUntil(until time.Time) {
	m.Called(until)
}
