package clock

import (
	"time"
)

type Clock interface {
	Now() time.Time
	Today() (year int, month time.Month, day int)
}

type realClock struct{}

func NewClock() *realClock {
	return &realClock{}
}

func (clock *realClock) Now() time.Time {
	return time.Now()
}

func (clock *realClock) Today() (year int, month time.Month, day int) {
	return time.Now().Date()
}

type Sleeper struct {
	clock   Clock
	sleepFn func(duration time.Duration)
}

func NewSleeper() *Sleeper {
	return &Sleeper{
		clock:   NewClock(),
		sleepFn: time.Sleep,
	}
}

func (sleeper *Sleeper) Sleep(duration time.Duration) {
	sleeper.sleepFn(duration)
}

func (sleeper *Sleeper) SleepUntil(until time.Time) {
	now := sleeper.clock.Now()
	time_to_sleep := until.Sub(now)

	if time_to_sleep > 0*time.Second {
		sleeper.sleepFn(time_to_sleep)
	}
}
