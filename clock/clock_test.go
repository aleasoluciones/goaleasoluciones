// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package clock_test

import (
	"testing"
	"time"

	. "github.com/aleasoluciones/goaleasoluciones/clock"

	"github.com/stretchr/testify/assert"
)

func TestClockReturnsCurrentTime(t *testing.T) {
	t.Parallel()
	clock := NewClock()

	result := clock.Now()

	assert.IsType(t, result, time.Time{})
}

func TestClockReturnsToday(t *testing.T) {
	t.Parallel()
	clock := NewClock()

	year, month, day := clock.Today()

	expected_year, expected_month, expected_day := time.Now().Date()
	assert.Equal(t, year, expected_year)
	assert.Equal(t, month, expected_month)
	assert.Equal(t, day, expected_day)
}

func TestSleeperSleep(t *testing.T) {
	t.Parallel()
	sleeper := NewSleeper()

	sleeper.Sleep(1 * time.Millisecond)
}

type FakeClock struct{}

func (clock *FakeClock) Now() time.Time {
	return time.Date(2013, time.September, 9, 23, 0, 0, 0, time.UTC)
}

func (clock *FakeClock) Today() (year int, month time.Month, day int) {
	return 2013, time.September, 9
}

func TestSleeperSleepUntilIfDurationIsAheadOfCurrentTime(t *testing.T) {
	t.Parallel()
	sleeper := NewSleeperWithClock(&FakeClock{})

	start := time.Now()
	sleeper.SleepUntil(time.Date(2013, time.September, 9, 23, 0, 1, 0, time.UTC))

	assert.WithinDuration(t, time.Now(), start, 2*time.Second)
}

func TestSleeperDoNotSleepUntilIfDurationIsBehindOfCurrentTime(t *testing.T) {
	t.Parallel()
	sleeper := NewSleeperWithClock(&FakeClock{})

	start := time.Now()
	sleeper.SleepUntil(time.Date(2013, time.September, 9, 22, 0, 0, 0, time.UTC))

	assert.WithinDuration(t, time.Now(), start, 1*time.Second)
}
