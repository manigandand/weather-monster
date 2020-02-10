package forecasts_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-monster/pkg/respond"
	"weather-monster/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forecasts API Test Suite", func() {
	Context("invalid city", func() {
		When("no cities", func() {
			var (
				res      *http.Response
				err      error
				response struct {
					Data *schema.Forecast `json:"data"`
					Meta respond.Meta     `json:"meta"`
				}
			)
			BeforeEach(func() {
			})
			JustBeforeEach(func() {
				res, err = tClient.GetForecast(12)
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should return empty response", func() {
				jsonErr := json.NewDecoder(res.Body).Decode(&response)
				Expect(jsonErr).ShouldNot(HaveOccurred())
				fmt.Println(response)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(response.Meta.Message).To(Equal("invalid city id"))
				Expect(response.Meta.Status).To(Equal(http.StatusBadRequest))
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

		Context("no temperature data", func() {
			When("no temperature data", func() {
				var (
					res      *http.Response
					err      error
					response struct {
						Data *schema.Forecast `json:"data"`
						Meta respond.Meta     `json:"meta"`
					}
				)
				BeforeEach(func() {
				})
				JustBeforeEach(func() {
					res, err = tClient.GetForecast(1)
					Expect(err).ShouldNot(HaveOccurred())
				})

				It("should return empty response", func() {
					jsonErr := json.NewDecoder(res.Body).Decode(&response)
					Expect(jsonErr).ShouldNot(HaveOccurred())
					fmt.Println(response)
					Expect(res.StatusCode).To(Equal(http.StatusOK))
					Expect(response.Data.CityID).To(Equal(uint(1)))
					Expect(response.Data.City.Name).To(Equal("Berlin"))
					Expect(response.Data.Min).To(Equal(float64(0)))
					Expect(response.Data.Max).To(Equal(float64(0)))
					Expect(response.Data.Sample).To(Equal(0))
					Expect(response.Meta.Status).To(Equal(http.StatusOK))
				})
			})
		})

		Context("temperature data", func() {
			When("save temperature with valid city", func() {
				var (
					req      []*schema.Temperature
					res      *http.Response
					err      error
					response struct {
						Data *schema.Temperature `json:"data"`
						Meta respond.Meta        `json:"meta"`
					}
				)

				BeforeEach(func() {
					req = []*schema.Temperature{
						&schema.Temperature{
							CityID: 1,
							Min:    10,
							Max:    25,
						},
						&schema.Temperature{
							CityID: 1,
							Min:    8,
							Max:    20,
						},
						&schema.Temperature{
							CityID: 1,
							Min:    14,
							Max:    29,
						},
						&schema.Temperature{
							CityID: 1,
							Min:    18,
							Max:    32,
						},
					}
				})

				It("should not fail", func() {
					for i, r := range req {
						res, err = tClient.PostTemperature(r)
						Expect(err).ShouldNot(HaveOccurred())

						jsonErr := json.NewDecoder(res.Body).Decode(&response)
						Expect(jsonErr).ShouldNot(HaveOccurred())
						fmt.Println(response)
						Expect(res.StatusCode).To(Equal(http.StatusCreated))
						Expect(response.Data.ID).To(Equal(uint(i + 1)))
						Expect(response.Meta.Status).To(Equal(http.StatusCreated))
					}
				})
			})

			When("no temperature data", func() {
				var (
					res      *http.Response
					err      error
					response struct {
						Data *schema.Forecast `json:"data"`
						Meta respond.Meta     `json:"meta"`
					}
				)
				BeforeEach(func() {
				})
				JustBeforeEach(func() {
					res, err = tClient.GetForecast(1)
					Expect(err).ShouldNot(HaveOccurred())
				})

				It("should return empty response", func() {
					jsonErr := json.NewDecoder(res.Body).Decode(&response)
					Expect(jsonErr).ShouldNot(HaveOccurred())
					fmt.Println(response)
					Expect(res.StatusCode).To(Equal(http.StatusOK))
					Expect(response.Data.CityID).To(Equal(uint(1)))
					Expect(response.Data.City.Name).To(Equal("Berlin"))
					Expect(response.Data.Min).To(Equal(float64(12.5)))
					Expect(response.Data.Max).To(Equal(float64(26.5)))
					Expect(response.Data.Sample).To(Equal(4))
					Expect(response.Meta.Status).To(Equal(http.StatusOK))
				})
			})
		})
	})
})
