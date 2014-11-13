package clock

import (
	"time"
)

type Clock interface {
	Now() time.Time
	Today() (year int, month time.Month, day int)
}

type clock struct{}

func NewClock() Clock {
	return &clock{}
}

func (c *clock) Now() time.Time {
	return time.Now()
}

func (c *clock) Today() (year int, month time.Month, day int) {
	return time.Now().Date()
}

type Sleeper interface {
	Sleep(duration time.Duration)
	SleepUntil(until time.Time)
}

type sleeper struct {
	clock Clock
}

func NewSleeper() Sleeper {
	return &sleeper{
		clock: NewClock(),
	}
}

func NewSleeperWithClock(clock Clock) Sleeper {
	return &sleeper{
		clock: clock,
	}
}

func (s *sleeper) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func (s *sleeper) SleepUntil(until time.Time) {
	now := s.clock.Now()
	time_to_sleep := until.Sub(now)

	if time_to_sleep > 0*time.Second {
		s.Sleep(time_to_sleep)
	}
}
