package temperatures_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-monster/pkg/respond"
	"weather-monster/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Temperature API test suite", func() {
	Context("no city registered", func() {
		When("no cities", func() {
			var (
				res      *http.Response
				err      error
				response struct {
					Data []*schema.Temperature `json:"data"`
					Meta respond.Meta          `json:"meta"`
				}
			)
			BeforeEach(func() {
			})
			JustBeforeEach(func() {
				res, err = tClient.GetTemperature()
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should return empty response", func() {
				jsonErr := json.NewDecoder(res.Body).Decode(&response)
				Expect(jsonErr).ShouldNot(HaveOccurred())
				fmt.Println(response)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(response.Data)).To(Equal(0))
				Expect(response.Meta.Status).To(Equal(http.StatusOK))
			})
		})
	})

	Context("valid city", func() {
		When("valid city req", func() {
			var (
				req      *schema.CityReq
				res      *http.Response
				err      error
				response struct {
					Data *schema.City `json:"data"`
					Meta respond.Meta `json:"meta"`
				}
			)

			BeforeEach(func() {
				req = &schema.CityReq{
					Name:      "Berlin",
					Latitude:  10.1111,
					Longitude: 32.2232,
				}
			})
			JustBeforeEach(func() {
				res, err = tClient.PostCities(req)
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should create a city", func() {
				jsonErr := json.NewDecoder(res.Body).Decode(&response)
				Expect(jsonErr).ShouldNot(HaveOccurred())
				fmt.Println(response)
				Expect(res.StatusCode).To(Equal(http.StatusCreated))
				Expect(response.Data.ID).To(Equal(uint(1)))
				Expect(response.Data.Name).To(Equal("Berlin"))
				Expect(response.Meta.Status).To(Equal(http.StatusCreated))
			})
		})

		When("save temperature with invalid city", func() {
			var (
				req      *schema.Temperature
				res      *http.Response
				err      error
				response struct {
					Data *schema.Temperature `json:"data"`
					Meta respond.Meta        `json:"meta"`
				}
			)

			BeforeEach(func() {
				req = &schema.Temperature{
					CityID: 122,
					Min:    10,
					Max:    32,
				}
			})
			JustBeforeEach(func() {
				res, err = tClient.PostTemperature(req)
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should fail", func() {
				jsonErr := json.NewDecoder(res.Body).Decode(&response)
				Expect(jsonErr).ShouldNot(HaveOccurred())
				fmt.Println(response)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(response.Meta.Message).To(Equal("invalid city id"))
				Expect(response.Meta.Status).To(Equal(http.StatusBadRequest))
			})
		})

		When("save temperature with valid city", func() {
			var (
				req      *schema.Temperature
				res      *http.Response
				err      error
				response struct {
					Data *schema.Temperature `json:"data"`
					Meta respond.Meta        `json:"meta"`
				}
			)

			BeforeEach(func() {
				req = &schema.Temperature{
					CityID: 1,
					Min:    10,
					Max:    32,
				}
			})
			JustBeforeEach(func() {
				res, err = tClient.PostTemperature(req)
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should fail", func() {
				jsonErr := json.NewDecoder(res.Body).Decode(&response)
				Expect(jsonErr).ShouldNot(HaveOccurred())
				fmt.Println(response)
				Expect(res.StatusCode).To(Equal(http.StatusCreated))
				Expect(response.Data.ID).To(Equal(uint(1)))
				Expect(response.Meta.Status).To(Equal(http.StatusCreated))
			})
		})

		Context("valid webhook registered with city", func() {
			When("save webhook for the valid city", func() {
				var (
					req      *schema.Webhook
					res      *http.Response
					err      error
					response struct {
						Data *schema.Temperature `json:"data"`
						Meta respond.Meta        `json:"meta"`
					}
				)

				BeforeEach(func() {
					req = &schema.Webhook{
						CityID:      1,
						CallbackURL: "http://webhook.io/aasassa",
					}
				})
				JustBeforeEach(func() {
					res, err = tClient.PostWebhooks(req)
					Expect(err).ShouldNot(HaveOccurred())
				})
				It("should fail", func() {
					jsonErr := json.NewDecoder(res.Body).Decode(&response)
					Expect(jsonErr).ShouldNot(HaveOccurred())
					fmt.Println(response)
					Expect(res.StatusCode).To(Equal(http.StatusCreated))
					Expect(response.Data.ID).To(Equal(uint(1)))
					Expect(response.Meta.Status).To(Equal(http.StatusCreated))
				})
			})

			When("save webhook for the valid city", func() {
				var (
					req      *schema.Webhook
					res      *http.Response
					err      error
					response struct {
						Data *schema.Temperature `json:"data"`
						Meta respond.Meta        `json:"meta"`
					}
				)

				BeforeEach(func() {
					req = &schema.Webhook{
						CityID:      1,
						CallbackURL: "https://webhook.site/3eebcb3e-b207-4aa5-90fa-601664498299",
					}
				})
				JustBeforeEach(func() {
					res, err = tClient.PostWebhooks(req)
					Expect(err).ShouldNot(HaveOccurred())
				})
				It("should fail", func() {
					jsonErr := json.NewDecoder(res.Body).Decode(&response)
					Expect(jsonErr).ShouldNot(HaveOccurred())
					fmt.Println(response)
					Expect(res.StatusCode).To(Equal(http.StatusCreated))
					Expect(response.Data.ID).To(Equal(uint(2)))
					Expect(response.Meta.Status).To(Equal(http.StatusCreated))
				})
			})

			When("save temperature with valid city", func() {
				var (
					req      *schema.Temperature
					res      *http.Response
					err      error
					response struct {
						Data *schema.Temperature `json:"data"`
						Meta respond.Meta        `json:"meta"`
					}
				)

				BeforeEach(func() {
					req = &schema.Temperature{
						CityID: 1,
						Min:    20,
						Max:    32,
					}
				})
				JustBeforeEach(func() {
					res, err = tClient.PostTemperature(req)
					Expect(err).ShouldNot(HaveOccurred())
				})
				It("should fail", func() {
					jsonErr := json.NewDecoder(res.Body).Decode(&response)
					Expect(jsonErr).ShouldNot(HaveOccurred())
					fmt.Println(response)
					Expect(res.StatusCode).To(Equal(http.StatusCreated))
					Expect(response.Data.ID).To(Equal(uint(2)))
					Expect(response.Meta.Status).To(Equal(http.StatusCreated))
				})
			})
		})
	})
})
