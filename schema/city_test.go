package schema_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "weather-monster/schema"
)

var _ = Describe("City schema test suite", func() {
	Context("City schema validate", func() {
		When("no name", func() {
			var (
				req *City
				err error
			)
			BeforeEach(func() {
				req = &City{
					Name: "",
				}
			})
			JustBeforeEach(func() {
				err = req.Ok()

			})
			It("should return error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(Equal("name is required"))
			})
		})

		When("try to update deleted field explicitly", func() {
			var (
				req *City
				err error
			)
			BeforeEach(func() {
				req = &City{
					Name:    "Berlin",
					Deleted: true,
				}
			})
			JustBeforeEach(func() {
				err = req.Ok()

			})
			It("should return error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(Equal("invalid request, you can't update deleted field"))
			})
		})

		When("Valid req", func() {
			var (
				req *City
				err error
			)
			BeforeEach(func() {
				req = &City{
					Name:    "Berlin",
					Deleted: false,
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

	Context("CityReq validate", func() {
		When("no name", func() {
			var (
				req *CityReq
				err error
			)
			BeforeEach(func() {
				req = &CityReq{
					Name: "",
				}
			})
			JustBeforeEach(func() {
				err = req.Ok()

			})
			It("should return error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(Equal("name is required"))
			})
		})

		When("Valid req", func() {
			var (
				req *CityReq
				err error
			)
			BeforeEach(func() {
				req = &CityReq{
					Name: "Berlin",
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
