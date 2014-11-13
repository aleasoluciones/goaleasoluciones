package tryto_test

import (
	"errors"

	"github.com/aleasoluciones/goaleasoluciones/mocks"
	. "github.com/aleasoluciones/goaleasoluciones/tryto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Try to", func() {
	var wasCalledTimes, aResult int

	BeforeEach(func() {
		wasCalledTimes = 0
		aResult = 42
	})

	Context("when wrapped func returns error", func() {
		It("retries the amount of times", func() {
			retrier := NewRetrier()

			retrier.RunRetrying(func() (interface{}, error) {
				wasCalledTimes++
				return 0, errors.New("an error")
			})

			Expect(wasCalledTimes).To(Equal(retrier.MaximumAttempts))
		})

		It("returns result in the wrapped function", func() {
			result, err := NewRetrier().RunRetrying(func() (interface{}, error) {
				wasCalledTimes++
				return aResult, errors.New("an error")
			})

			Expect(result.(int)).To(Equal(aResult))
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when wrapped func is successful at first attempt", func() {
		It("tries only once", func() {
			NewRetrier().RunRetrying(func() (interface{}, error) {
				wasCalledTimes++
				return 0, nil
			})

			Expect(wasCalledTimes).To(Equal(1))
		})
	})

	Context("when wrapped func is successful after some tries", func() {
		It("retries several times", func() {
			NewRetrier().RunRetrying(func() (interface{}, error) {
				if wasCalledTimes == 2 {
					return aResult, nil
				}
				wasCalledTimes++
				return 0, errors.New("an error")
			})

			Expect(wasCalledTimes).To(Equal(2))
		})

		It("returns last result", func() {
			result, err := NewRetrier().RunRetrying(func() (interface{}, error) {
				if wasCalledTimes == 2 {
					return aResult, nil
				}
				wasCalledTimes++
				return 0, errors.New("an error")
			})

			Expect(result.(int)).To(Equal(aResult))
			Expect(err).ToNot(HaveOccurred())
		})
	})

	FContext("when wrapped func returns error and lablabla", func() {
		It("sleep FOO", func() {
			sleeper := new(mocks.Sleeper)
			retrier := NewRetrierWithSleeper(sleeper)
			sleeper.On("Sleep", retrier.Interval).Return(nil)

			retrier.RunRetrying(func() (interface{}, error) {
				wasCalledTimes++
				return 0, errors.New("an error")
			})

			sleeper.AssertNumberOfCalls(GinkgoT(), "Sleep", retrier.MaximumAttempts-1)
		})
	})
})
