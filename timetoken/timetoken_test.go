// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package timetoken_test

import (
	"testing"
	"time"

	. "github.com/aleasoluciones/goaleasoluciones/timetoken"

	"github.com/stretchr/testify/assert"
)

const (
	ZERO_TTL_FOR_TOKEN_MANAGER  = 0 * time.Millisecond
	ONE_CYCLE_FOR_TOKEN_MANAGER = 700 * time.Millisecond
)

func tokenManager(periodicFuncTimesCalled *int) *TokenManager {
	return NewTokenManager(ONE_CYCLE_FOR_TOKEN_MANAGER, ZERO_TTL_FOR_TOKEN_MANAGER, func(id string) {
		*periodicFuncTimesCalled += 1
	})
}

func TestPeriodicFunctionCalledUntilTokenTTL(t *testing.T) {
	t.Parallel()
	periodicFuncTimesCalled := 0
	tm := tokenManager(&periodicFuncTimesCalled)
	var tokenTTL = 2 * ONE_CYCLE_FOR_TOKEN_MANAGER
	tm.Add("id", tokenTTL)
	time.Sleep(tokenTTL + 50*time.Millisecond)
	assert.Equal(t, 2, periodicFuncTimesCalled)
}

func TestPeriodicFunctionCalledEveryPeriodeUntilTokenTTL(t *testing.T) {
	t.Parallel()
	periodicFuncTimesCalled := 0
	tm := tokenManager(&periodicFuncTimesCalled)
	var tokenTTL = 3 * ONE_CYCLE_FOR_TOKEN_MANAGER
	var lessThanOneCycle = ONE_CYCLE_FOR_TOKEN_MANAGER - 100*time.Millisecond
	tm.Add("id", tokenTTL)
	time.Sleep(lessThanOneCycle)
	assert.Equal(t, 1, periodicFuncTimesCalled)
	time.Sleep(lessThanOneCycle)
	assert.Equal(t, 2, periodicFuncTimesCalled)
	time.Sleep(lessThanOneCycle)
	assert.Equal(t, 3, periodicFuncTimesCalled)
	time.Sleep(lessThanOneCycle)
	assert.Equal(t, 3, periodicFuncTimesCalled)
}

func TestPeriodicFunctionNotCalledAfterTokenTTL(t *testing.T) {
	t.Parallel()
	periodicFuncTimesCalled := 0
	tm := tokenManager(&periodicFuncTimesCalled)
	var tokenTTL = 1 * ONE_CYCLE_FOR_TOKEN_MANAGER
	tm.Add("id", tokenTTL)
	time.Sleep(2 * tokenTTL)
	assert.Equal(t, 1, periodicFuncTimesCalled)
}
