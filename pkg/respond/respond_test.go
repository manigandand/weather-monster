package respond_test

/*
import (
	"net/http"
	"net/http/httptest"
	"weather-monster/pkg/errors"
	. "weather-monster/pkg/respond"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Respond", func() {
	var testData = map[string]interface{}{"test": true}
	It("Test With - should respond custom data", func() {
		w := httptest.NewRecorder()
		OK(w, testData)

		Expect(w.Code).To(Equal(http.StatusOK))
		Expect(w.HeaderMap.Get("Content-Type")).To(Equal("application/json"))
		Expect(w.HeaderMap.Get("Content-Encoding")).To(Equal("gzip"))
	})

	It("New Page - should decode paginate information from request", func() {
		r, err := http.NewRequest("GET", "recipe", nil)
		Ω(err).ShouldNot(HaveOccurred())
		urls := r.URL.Query()
		urls.Add("limit", "10")
		urls.Add("offset", "10")
		r.URL.RawQuery = urls.Encode()
		r.ParseForm()
	})

	It("New Page - max limit should be 10. server should reset limit to 10", func() {
		r, err := http.NewRequest("GET", "recipe", nil)
		urls := r.URL.Query()
		urls.Add("limit", "20")
		urls.Add("offset", "10")
		r.URL.RawQuery = urls.Encode()
		r.ParseForm()

		Ω(err).ShouldNot(HaveOccurred())
	})
	It("Foramt - should respond formated response", func() {
		w := httptest.NewRecorder()
		Format(w, http.StatusOK, testData)
		Expect(w.Code).To(Equal(http.StatusOK))
		Expect(w.HeaderMap.Get("Content-Type")).To(Equal("application/json"))
		Expect(w.HeaderMap.Get("Content-Encoding")).To(Equal("gzip"))
	})
	It("Fail - should return error response", func() {
		w := httptest.NewRecorder()
		Fail(w, errors.BadRequest("invalid recipe id"))
		Expect(w.Code).To(Equal(http.StatusBadRequest))
		Expect(w.HeaderMap.Get("Content-Type")).To(Equal("application/json"))
		Expect(w.HeaderMap.Get("Content-Encoding")).To(Equal("gzip"))
	})
})
*/
