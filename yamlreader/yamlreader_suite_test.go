// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package yamlreader

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestYamlReader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "YamlReader Suite")
}
