package retrier_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTryto(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Retrier Suite")
}
