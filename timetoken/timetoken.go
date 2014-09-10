package timetoken

import (
	"sync"
	"time"

	"github.com/aleasoluciones/goaleasoluciones/scheduledtask"
)

type TokenManager struct {
	periode      time.Duration
	ttl          time.Duration
	periodicFunc func(id string)
	tokens       map[string]time.Time
	mutex        sync.Mutex
}

func New(periode, ttl time.Duration, periodicFunc func(id string)) *TokenManager {
	tm := TokenManager{
		periode:      periode,
		ttl:          ttl,
		periodicFunc: periodicFunc,
		tokens:       make(map[string]time.Time),
	}

	scheduledtask.NewScheduledTask(
		tm.executePeriodicFunc,
		tm.periode,
		tm.ttl)
	return &tm
}

func (tm TokenManager) executePeriodicFunc() {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	now := time.Now()

	for k, v := range tm.tokens {
		if v.Sub(now) >= 0 {
			go tm.periodicFunc(k)
		} else {
			delete(tm.tokens, k)
		}
	}
}

func (tm *TokenManager) Add(id string, ttl time.Duration) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	newExpirationTime := time.Now().Add(ttl)

	actualExpirationTime, found := tm.tokens[id]
	if !found {
		tm.tokens[id] = newExpirationTime
		return
	}

	if newExpirationTime.Sub(actualExpirationTime) > 0 {
		tm.tokens[id] = newExpirationTime
	}
}
