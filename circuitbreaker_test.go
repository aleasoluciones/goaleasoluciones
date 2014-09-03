package gocircuitbreaker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TODO
// si mandamos tantos errores como numErrores configurados el circuit IsOpen
// si mandamos un ok (despues de que este abierto), se cierra IsClosed
// si esperamos el tiempo de reset despues de IsOpen se cierra IsClosed
// si mandamos tantos errores como numErrores -1 y esperamos el timpo de reset, aunque enviemos otro
// error, el circuito sigue cerrado IsClosed

func TestIsOpenWhenItHasTheNumberOfConfiguredErrors(t *testing.T) {
	assert := assert.New(t)
	circuit := NewCircuit(3, time.Duration(5)*time.Second)
	circuit.Error()
	circuit.Error()
	circuit.Error()
	assert.Equal(circuit.IsOpen(), true, "the circuit should be open")
}
