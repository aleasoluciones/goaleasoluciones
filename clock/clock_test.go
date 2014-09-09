package clock

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)


func TestClockReturnsCurrentTime(t *testing.T) {
	clock := NewClock()

	result := clock.Now()

	assert.IsType(t, result, time.Time{})
}

func TestClockReturnsToday(t *testing.T) {
	clock := NewClock()

	year, month, day := clock.Today()

	expected_year, expected_month, expected_day := time.Now().Date()
	assert.Equal(t, year, expected_year)
	assert.Equal(t, month, expected_month)
	assert.Equal(t, day, expected_day)
}

func TestSleeperSleep(t *testing.T) {
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
	sleeper := NewSleeper()
	clock := &FakeClock{}
	sleeper.clock = clock
	var sleepFnCalled time.Duration
	sleeper.sleepFn = func(duration time.Duration) { sleepFnCalled = duration }

	sleeper.SleepUntil(time.Date(2013, time.September, 9, 23, 0, 5, 0, time.UTC))

	assert.Equal(t, sleepFnCalled, 5 * time.Second)
}

func TestSleeperDoNotSleepUntilIfDurationIsBehindOfCurrentTime(t *testing.T) {
	sleeper := NewSleeper()
	clock := &FakeClock{}
	sleeper.clock = clock
	sleepFnCalled := false
	sleeper.sleepFn = func(duration time.Duration) { sleepFnCalled = true }

	sleeper.SleepUntil(time.Date(2013, time.September, 9, 22, 0, 0, 0, time.UTC))

	assert.False(t, sleepFnCalled)
}
