// Copyright 2015 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package periodictask_test

import (
	"testing"
	"time"

	"github.com/aleasoluciones/goaleasoluciones/periodictask"

	"github.com/stretchr/testify/assert"
)

func TestExecutionEachSecond(t *testing.T) {
	t.Parallel()

	var counter int = 0
	counterFunc := func() {
		counter = counter + 1
	}

	periodictask.New(counterFunc, "* * * * * * *").Run()
	time.Sleep(3 * time.Second)
	assert.True(t, counter >= 2)
	assert.True(t, counter <= 3)
}
