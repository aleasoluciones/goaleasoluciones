// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package circuitbreaker_test

import (
	"testing"
	"time"

	. "github.com/aleasoluciones/goaleasoluciones/circuitbreaker"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_TICK = 200 * time.Microsecond
)

func newCircuit() *Circuit {
	return NewCircuit(3, TEST_TICK)
}

func TestOpenWhenItHasTheNumberOfConfiguredErrors(t *testing.T) {
	t.Parallel()
	circuit := newCircuit()
	circuit.Error()
	circuit.Error()
	circuit.Error()
	assert.True(t, circuit.IsOpen())
}

func TestClosedAfterAnOKWhenItWasOpen(t *testing.T) {
	t.Parallel()
	circuit := newCircuit()
	circuit.Error()
	circuit.Error()
	circuit.Error()
	waitUntilReset()
	circuit.Ok()
	assert.True(t, circuit.IsClosed())
}

func TestClosedWhenItIsOpenAndResetTimeHasExpired(t *testing.T) {
	t.Parallel()
	circuit := newCircuit()
	circuit.Error()
	circuit.Error()
	circuit.Error()
	waitUntilReset()
	circuit.Ok()
	assert.True(t, circuit.IsClosed())
}

func TestClosedWhenTheCircuitHasBeenResetAsTimeExpired(t *testing.T) {
	t.Parallel()
	circuit := newCircuit()
	circuit.Error()
	circuit.Error()
	waitUntilReset()
	circuit.Error()
	assert.True(t, circuit.IsClosed())
}

func waitUntilReset() {
	time.Sleep(TEST_TICK)
	time.Sleep(50 * time.Microsecond)
}
