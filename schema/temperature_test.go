package schema_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "weather-monster/schema"
)

var _ = Describe("Temperature schema test suite", func() {
	Context("Temperature schema validate", func() {
		When("no city id", func() {
			var (
				req *Temperature
				err error
			)
			BeforeEach(func() {
				req = &Temperature{
					CityID: 0,
				}
			})
			JustBeforeEach(func() {
				err = req.Ok()
			})
			It("should return error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(Equal("city id is required"))
			})
		})

		When("valid req", func() {
			var (
				req *Temperature
				err error
			)
			BeforeEach(func() {
				req = &Temperature{
					CityID: 2,
				}
			})
			JustBeforeEach(func() {
				err = req.Ok()
			})
			It("should return nil", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
