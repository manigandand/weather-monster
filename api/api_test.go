package api_test

import (
	"encoding/json"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "weather-monster/api"
)

var _ = Describe("Api", func() {
	Context("Basic api http://127.0.0.1:8080/ test", func() {
		It("GET index", func() {
			InitService("test", "1")

			res, err := tClient.WeatherMonsterHomePage()
			Ω(err).ShouldNot(HaveOccurred())
			var data interface{}
			jsonErr := json.NewDecoder(res.Body).Decode(&data)
			Ω(jsonErr).ShouldNot(HaveOccurred())
			// fmt.Println(data)
			Expect(res.StatusCode).To(Equal(http.StatusOK))
		})

		It("GET /top handler", func() {

			res, err := tClient.WeatherMonsterTopPage()
			Ω(err).ShouldNot(HaveOccurred())
			var data interface{}
			jsonErr := json.NewDecoder(res.Body).Decode(&data)
			Ω(jsonErr).ShouldNot(HaveOccurred())
			// fmt.Println(data)
			Expect(res.StatusCode).To(Equal(http.StatusOK))
		})
	})
})

func (c *Client) WeatherMonsterHomePage() (*http.Response, error) {
	return c.DoGet("/")
}

func (c *Client) WeatherMonsterTopPage() (*http.Response, error) {
	return c.DoGet("/top")
}
