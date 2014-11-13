package tryto_test

import (
	"errors"

	. "github.com/aleasoluciones/goaleasoluciones/tryto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Try to", func() {
	var wasCalledTimes, maximumAttempts, aResult int

	BeforeEach(func() {
		wasCalledTimes = 0
		maximumAttempts = 5
		aResult = 42
	})

	Context("when wrapped func returns error", func() {
		It("retries the amount of times", func() {
			TryTo(func() (interface{}, error) {
				wasCalledTimes++
				return 0, errors.New("an error")
			}, maximumAttempts)

			Expect(wasCalledTimes).To(Equal(maximumAttempts))
		})

		It("returns result in the wrapped function", func() {
			result, err := TryTo(func() (interface{}, error) {
				wasCalledTimes++
				return aResult, errors.New("an error")
			}, maximumAttempts)

			Expect(result.(int)).To(Equal(aResult))
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when wrapped func is successful at first attempt", func() {
		It("tries only once", func() {
			TryTo(func() (interface{}, error) {
				wasCalledTimes++
				return 0, nil
			}, maximumAttempts)

			Expect(wasCalledTimes).To(Equal(1))
		})
	})

	Context("when wrapped func is successful after some tries", func() {
		It("retries several times", func() {
			TryTo(func() (interface{}, error) {
				if wasCalledTimes == 2 {
					return aResult, nil
				}
				wasCalledTimes++
				return 0, errors.New("an error")
			}, maximumAttempts)

			Expect(wasCalledTimes).To(Equal(2))
		})

		It("returns last result", func() {
			result, err := TryTo(func() (interface{}, error) {
				if wasCalledTimes == 2 {
					return aResult, nil
				}
				wasCalledTimes++
				return 0, errors.New("an error")
			}, maximumAttempts)

			Expect(result.(int)).To(Equal(aResult))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
