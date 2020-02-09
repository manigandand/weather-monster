package v1

import (
	"weather-monster/api/v1/cities"
	"weather-monster/api/v1/forecasts"
	"weather-monster/api/v1/temperatures"
	"weather-monster/api/v1/webhooks"

	"github.com/go-chi/chi"
)

// Init initializes all the v1 routes
func Init(r chi.Router) {
	r.Route("/cities", cities.Init)
	r.Route("/temperatures", temperatures.Init)
	r.Route("/forecasts", forecasts.Init)
	r.Route("/webhooks", webhooks.Init)
}
