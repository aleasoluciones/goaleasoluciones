// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package retrier_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRetrier(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Retrier Suite")
}
