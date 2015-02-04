// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package retrier_test

import (
	"errors"

	"github.com/aleasoluciones/goaleasoluciones/mocks"
	. "github.com/aleasoluciones/goaleasoluciones/retrier"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Retrier", func() {
	var (
		wasCalledTimes, aResult int
		sleeper                 *mocks.Sleeper
		retrier                 *Retrier
	)

	BeforeEach(func() {
		wasCalledTimes = 0
		aResult = 42
		sleeper = new(mocks.Sleeper)
		retrier = NewRetrierWithSleeper(sleeper)
		sleeper.On("Sleep", retrier.Interval).Return(nil)
	})

	Context("when wrapped func returns error", func() {
		It("retries the amount of times", func() {
			retrier.RunRetrying(func() (interface{}, error) {
				wasCalledTimes++
				return 0, errors.New("an error")
			})

			Expect(wasCalledTimes).To(Equal(retrier.MaximumAttempts))
		})

		It("returns result in the wrapped function", func() {
			result, err := retrier.RunRetrying(func() (interface{}, error) {
				wasCalledTimes++
				return aResult, errors.New("an error")
			})

			Expect(result.(int)).To(Equal(aResult))
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when wrapped func is successful at first attempt", func() {
		It("tries only once", func() {
			retrier.RunRetrying(func() (interface{}, error) {
				wasCalledTimes++
				return 0, nil
			})

			Expect(wasCalledTimes).To(Equal(1))
		})
	})

	Context("when wrapped func is successful after some tries", func() {
		It("retries several times", func() {
			retrier.RunRetrying(func() (interface{}, error) {
				if wasCalledTimes == 2 {
					return aResult, nil
				}
				wasCalledTimes++
				return 0, errors.New("an error")
			})

			Expect(wasCalledTimes).To(Equal(2))
		})

		It("returns last result", func() {
			result, err := retrier.RunRetrying(func() (interface{}, error) {
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

	Context("when retrier support sleeper", func() {
		Context("when wrapped func returns error", func() {
			It("retries after period of time", func() {
				retrier.RunRetrying(func() (interface{}, error) {
					wasCalledTimes++
					return 0, errors.New("an error")
				})

				sleeper.AssertNumberOfCalls(GinkgoT(), "Sleep", retrier.MaximumAttempts-1)
			})
		})

		Context("when wrapped func is successful at first attempt", func() {
			It("do not use sleeper", func() {
				retrier.RunRetrying(func() (interface{}, error) {
					wasCalledTimes++
					return 0, errors.New("an error")
				})

				sleeper.AssertNotCalled(GinkgoT(), "Sleep")
			})
		})
	})
})
