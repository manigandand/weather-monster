package v1_test

import (
	"github.com/go-chi/chi"
	. "github.com/onsi/ginkgo"

	. "weather-monster/api/v1"
)

var _ = Describe("Api", func() {
	It("Init Routes", func() {
		rr := chi.NewRouter()
		rr.Route("/", Routes)
		Routes(rr)
	})
})
