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
