package cities_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-monster/pkg/respond"
	"weather-monster/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("City", func() {
	Context("GET /v1/cities", func() {
		When("no cities", func() {
			var (
				res      *http.Response
				err      error
				response struct {
					Data []*schema.City `json:"data"`
					Meta respond.Meta   `json:"meta"`
				}
			)
			BeforeEach(func() {
			})
			JustBeforeEach(func() {
				res, err = tClient.GetCities()
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

	Context("POST /v1/cities", func() {
		When("no city name", func() {
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
					Name:      "",
					Latitude:  10.1111,
					Longitude: 32.2232,
				}
			})
			JustBeforeEach(func() {
				res, err = tClient.PostCities(req)
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should fail", func() {
				jsonErr := json.NewDecoder(res.Body).Decode(&response)
				Expect(jsonErr).ShouldNot(HaveOccurred())
				fmt.Println(response)
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(response.Data).To(BeNil())
				Expect(response.Meta.Status).To(Equal(http.StatusBadRequest))
				Expect(response.Meta.Message).To(Equal("name is required"))
			})
		})
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
			It("should fail", func() {
				jsonErr := json.NewDecoder(res.Body).Decode(&response)
				Expect(jsonErr).ShouldNot(HaveOccurred())
				fmt.Println(response)
				Expect(res.StatusCode).To(Equal(http.StatusCreated))
				Expect(response.Data.ID).To(Equal(uint(1)))
				Expect(response.Data.Name).To(Equal("Berlin"))
				Expect(response.Meta.Status).To(Equal(http.StatusCreated))
			})
		})
		When("cities present", func() {
			var (
				res      *http.Response
				err      error
				response struct {
					Data []*schema.City `json:"data"`
					Meta respond.Meta   `json:"meta"`
				}
			)
			BeforeEach(func() {
			})
			JustBeforeEach(func() {
				res, err = tClient.GetCities()
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should return 1 city response", func() {
				jsonErr := json.NewDecoder(res.Body).Decode(&response)
				Expect(jsonErr).ShouldNot(HaveOccurred())
				fmt.Println(response)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(len(response.Data)).To(Equal(1))
				Expect(response.Meta.Status).To(Equal(http.StatusOK))
			})
		})

		Context("GET /v1/cities/:cityID", func() {
			When("get city by id", func() {
				var (
					res      *http.Response
					err      error
					response struct {
						Data *schema.City `json:"data"`
						Meta respond.Meta `json:"meta"`
					}
				)
				BeforeEach(func() {
				})
				JustBeforeEach(func() {
					res, err = tClient.GetCitiesByID(1)
					Expect(err).ShouldNot(HaveOccurred())
				})
				It("should return 1 city response", func() {
					jsonErr := json.NewDecoder(res.Body).Decode(&response)
					Expect(jsonErr).ShouldNot(HaveOccurred())
					fmt.Println(response)
					Expect(res.StatusCode).To(Equal(http.StatusOK))
					Expect(response.Data.ID).To(Equal(uint(1)))
					Expect(response.Data.Name).To(Equal("Berlin"))
					Expect(response.Meta.Status).To(Equal(http.StatusOK))
				})
			})
		})

		Context("PATCH /v1/cities/:cityID", func() {
			When("update city by id", func() {
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
						Name:      "Berlin-2",
						Latitude:  10.1111,
						Longitude: 32.2232,
					}
				})
				JustBeforeEach(func() {
					res, err = tClient.UpdateCities(1, req)
					Expect(err).ShouldNot(HaveOccurred())
				})
				It("should return 1 city response", func() {
					jsonErr := json.NewDecoder(res.Body).Decode(&response)
					Expect(jsonErr).ShouldNot(HaveOccurred())
					fmt.Println(response)
					Expect(res.StatusCode).To(Equal(http.StatusOK))
					Expect(response.Data.ID).To(Equal(uint(1)))
					Expect(response.Data.Name).To(Equal("Berlin-2"))
					Expect(response.Meta.Status).To(Equal(http.StatusOK))
				})
			})
		})

		Context("DELETE /v1/cities/:cityID", func() {
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
						Name:      "Bangalore",
						Latitude:  10.1111,
						Longitude: 32.2232,
					}
				})
				JustBeforeEach(func() {
					res, err = tClient.PostCities(req)
					Expect(err).ShouldNot(HaveOccurred())
				})
				It("should fail", func() {
					jsonErr := json.NewDecoder(res.Body).Decode(&response)
					Expect(jsonErr).ShouldNot(HaveOccurred())
					fmt.Println(response)
					Expect(res.StatusCode).To(Equal(http.StatusCreated))
					Expect(response.Data.ID).To(Equal(uint(2)))
					Expect(response.Data.Name).To(Equal("Bangalore"))
					Expect(response.Meta.Status).To(Equal(http.StatusCreated))
				})
			})

			When("delete city by id", func() {
				var (
					res *http.Response
					err error
				)
				BeforeEach(func() {
				})
				JustBeforeEach(func() {
					res, err = tClient.DeleteCity(2)
					Expect(err).ShouldNot(HaveOccurred())
				})
				It("should return 1 city response", func() {
					Expect(res.StatusCode).To(Equal(http.StatusNoContent))
				})
			})
		})

		Context("Extra routes /v1/cities/:cityID/*", func() {
			When("GET /v1/cities/:cityID/temperature", func() {
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
					res, err = tClient.GetCityTemperatureByID(1)
					Expect(err).ShouldNot(HaveOccurred())
				})
				It("should return empty city temperature response", func() {
					jsonErr := json.NewDecoder(res.Body).Decode(&response)
					Expect(jsonErr).ShouldNot(HaveOccurred())
					fmt.Println(response)
					Expect(res.StatusCode).To(Equal(http.StatusOK))
					Expect(len(response.Data)).To(Equal(0))
					Expect(response.Meta.Status).To(Equal(http.StatusOK))
				})
			})

			When("GET /v1/cities/:cityID/webhooks", func() {
				var (
					res      *http.Response
					err      error
					response struct {
						Data []*schema.Webhook `json:"data"`
						Meta respond.Meta      `json:"meta"`
					}
				)
				BeforeEach(func() {
				})
				JustBeforeEach(func() {
					res, err = tClient.GetCityWebhookByID(1)
					Expect(err).ShouldNot(HaveOccurred())
				})
				It("should return empty city webhooks response", func() {
					jsonErr := json.NewDecoder(res.Body).Decode(&response)
					Expect(jsonErr).ShouldNot(HaveOccurred())
					fmt.Println(response)
					Expect(res.StatusCode).To(Equal(http.StatusOK))
					Expect(len(response.Data)).To(Equal(0))
					Expect(response.Meta.Status).To(Equal(http.StatusOK))
				})
			})
		})
	})
})
