package schema_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "weather-monster/schema"
)

var _ = Describe("Webhook schema test suite", func() {
	Context("Webhook schema validate", func() {
		When("no city id", func() {
			var (
				req *Webhook
				err error
			)
			BeforeEach(func() {
				req = &Webhook{
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

		When("no call back url", func() {
			var (
				req *Webhook
				err error
			)
			BeforeEach(func() {
				req = &Webhook{
					CityID:      1,
					CallbackURL: "",
				}
			})
			JustBeforeEach(func() {
				err = req.Ok()

			})
			It("should return error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(Equal("call back url is required"))
			})
		})

		When("valid req", func() {
			var (
				req *Webhook
				err error
			)
			BeforeEach(func() {
				req = &Webhook{
					CityID:      1,
					CallbackURL: "https://ads.com",
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
