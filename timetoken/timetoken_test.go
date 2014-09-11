package timetoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TokenManager(periodicFuncTimesCalled *int) *tokenManager {
	return NewTokenManager(1*time.Millisecond, 3*time.Millisecond, func(id string) {
		*periodicFuncTimesCalled += 1
	})
}

func TestPeriodicFunctionCalledUntilTTL(t *testing.T) {
	t.Parallel()
	periodicFuncTimesCalled := 0
	tm := TokenManager(&periodicFuncTimesCalled)
	tm.Add("id", 3*time.Millisecond)
	time.Sleep(2 * time.Millisecond)
	assert.Equal(t, 2, periodicFuncTimesCalled)
}

func TestPeriodicFunctionCalledEveryPeriodeUntilTTL(t *testing.T) {
	t.Parallel()
	periodicFuncTimesCalled := 0
	tm := TokenManager(&periodicFuncTimesCalled)
	tm.Add("id", 3*time.Millisecond)

	time.Sleep(1 * time.Millisecond)
	assert.Equal(t, 1, periodicFuncTimesCalled)

	time.Sleep(1 * time.Millisecond)
	assert.Equal(t, 2, periodicFuncTimesCalled)
}

func TestPeriodicFunctionNotCalledAfterTTL(t *testing.T) {
	t.Parallel()
	periodicFuncTimesCalled := 0
	tm := TokenManager(&periodicFuncTimesCalled)
	tm.Add("id", 2*time.Millisecond)

	time.Sleep(3 * time.Millisecond)
	assert.Equal(t, 2, periodicFuncTimesCalled)
}
