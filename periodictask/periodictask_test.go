// Copyright 2015 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package periodictask_test

import (
	"testing"
	"time"

	"github.com/facebookgo/clock"

	. "github.com/aleasoluciones/goaleasoluciones/periodictask"
)

func TestSkeleton(t *testing.T) {
	t.Parallel()

	var counter int = 0
	counterFunc := func() {
		counter = counter + 1
	}

	mockClock := clock.NewMock()
	NewWithClock(counterFunc, "* * * * * *", mockClock)
	mockClock.Add(2 * time.Hour)
}
