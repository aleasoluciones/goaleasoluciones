package yamlreader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aleasoluciones/goaleasoluciones/yamlreader"
)

var _ = Describe("Yaml Reader", func() {
	yamlReader := NewYamlReader("../test_conf/test.yaml")

	Describe("Getting integer value", func() {
		Context("when value exists", func() {
			It("returns the value", func() {
				defaultValue := 0
				value, exists := yamlReader.GetIntValue("an_existing_integer_value", defaultValue)

				expectedValue := 42
				Expect(exists).To(Equal(true))
				Expect(value).To(Equal(expectedValue))
			})
		})
		Context("when value DOS NOT exist", func() {
			It("returns default value", func() {
				defaultValue := 0
				value, exists := yamlReader.GetIntValue("a_non_existing_value", defaultValue)

				Expect(exists).To(Equal(false))
				Expect(value).To(Equal(defaultValue))
			})
		})
		Context("when value exists but is not an integer", func() {
			It("returns default value", func() {
				defaultValue := -1
				value, exists := yamlReader.GetIntValue("an_existing_not_integer_value", defaultValue)

				Expect(exists).To(Equal(false))
				Expect(value).To(Equal(defaultValue))
			})
		})
	})

	Describe("Getting boolean value", func() {
		Context("when value exists", func() {
			It("returns the value", func() {
				defaultValue := false
				value, exists := yamlReader.GetBoolValue("an_existing_boolean_value", defaultValue)

				expectedValue := true
				Expect(exists).To(Equal(true))
				Expect(value).To(Equal(expectedValue))
			})
		})
		Context("when value DOS NOT exist", func() {
			It("returns default value", func() {
				defaultValue := false
				value, exists := yamlReader.GetBoolValue("a_non_existing_value", defaultValue)

				Expect(exists).To(Equal(false))
				Expect(value).To(Equal(defaultValue))
			})
		})
		Context("when value exists but is not a boolean", func() {
			It("returns default value", func() {
				defaultValue := false
				value, exists := yamlReader.GetBoolValue("an_existing_not_boolean_value", defaultValue)

				Expect(exists).To(Equal(false))
				Expect(value).To(Equal(defaultValue))
			})
		})
	})

})
