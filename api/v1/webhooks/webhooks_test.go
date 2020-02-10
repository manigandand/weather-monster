package webhooks_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-monster/pkg/respond"
	"weather-monster/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Webhooks API test suite", func() {
	Context("no city registered", func() {
		When("no cities", func() {
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
				res, err = tClient.GetWebhook()
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

		When("save webhook with invalid city", func() {
			var (
				req      *schema.Webhook
				res      *http.Response
				err      error
				response struct {
					Data *schema.Webhook `json:"data"`
					Meta respond.Meta    `json:"meta"`
				}
			)

			BeforeEach(func() {
				req = &schema.Webhook{
					CityID:      100,
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
				Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(response.Meta.Message).To(Equal("invalid city id"))
				Expect(response.Meta.Status).To(Equal(http.StatusBadRequest))
			})
		})

		When("save webhook for the valid city", func() {
			var (
				req      *schema.Webhook
				res      *http.Response
				err      error
				response struct {
					Data *schema.Webhook `json:"data"`
					Meta respond.Meta    `json:"meta"`
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
			It("should not fail", func() {
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
					Data *schema.Webhook `json:"data"`
					Meta respond.Meta    `json:"meta"`
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
			It("should not fail", func() {
				jsonErr := json.NewDecoder(res.Body).Decode(&response)
				Expect(jsonErr).ShouldNot(HaveOccurred())
				fmt.Println(response)
				Expect(res.StatusCode).To(Equal(http.StatusCreated))
				Expect(response.Data.ID).To(Equal(uint(2)))
				Expect(response.Meta.Status).To(Equal(http.StatusCreated))
			})
		})

	})

	Context("GET /v1/webhooks/:webhookID", func() {
		When("get webhook by id", func() {
			var (
				res      *http.Response
				err      error
				response struct {
					Data *schema.Webhook `json:"data"`
					Meta respond.Meta    `json:"meta"`
				}
			)
			BeforeEach(func() {
			})
			JustBeforeEach(func() {
				res, err = tClient.GetWebhookByID(1)
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should return 1 city response", func() {
				jsonErr := json.NewDecoder(res.Body).Decode(&response)
				Expect(jsonErr).ShouldNot(HaveOccurred())
				fmt.Println(response)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(response.Data.ID).To(Equal(uint(1)))
				Expect(response.Meta.Status).To(Equal(http.StatusOK))
			})
		})
	})

	Context("DELETE /v1/webhooks/:webhookID", func() {
		When("delete city by id", func() {
			var (
				res *http.Response
				err error
			)
			BeforeEach(func() {
			})
			JustBeforeEach(func() {
				res, err = tClient.DeleteWebhook(2)
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should return 1 city response", func() {
				Expect(res.StatusCode).To(Equal(http.StatusNoContent))
			})
		})
	})

})
