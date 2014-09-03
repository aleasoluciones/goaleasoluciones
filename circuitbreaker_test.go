package circuitbreaker

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

func newCircuit() *Circuit {
    return NewCircuit(3, 1*time.Second)
}

func TestOpenWhenItHasTheNumberOfConfiguredErrors(t *testing.T) {
    circuit := newCircuit()
    circuit.Error()
    circuit.Error()
    circuit.Error()
    assert.Equal(t, circuit.IsOpen(), true)
}

func TestClosedAfterAnOKWhenItWasOpen(t *testing.T) {
    circuit := newCircuit()
    circuit.Error()
    circuit.Error()
    circuit.Error()
    waitUntilReset()
    circuit.Ok()
    assert.Equal(t, circuit.IsClosed(), true)
}

func TestClosedWhenItIsOpenAndResetTimeHasExpired(t *testing.T) {
    circuit := newCircuit()
    circuit.Error()
    circuit.Error()
    circuit.Error()
    waitUntilReset()
    circuit.Ok()
    assert.Equal(t, circuit.IsClosed(), true)
}

func TestClosedWhenTheCircuitHasBeenResetAsTimeExpired(t *testing.T) {
    circuit := newCircuit()
    circuit.Error()
    circuit.Error()
    waitUntilReset()
    circuit.Error()
    assert.Equal(t, circuit.IsClosed(), true)
}

func waitUntilReset() {
    time.Sleep(1 * time.Second)
}
