package v1

import (
	"fmt"
	"net/http"
	"weather-monster/api"
	"weather-monster/api/v1/cities"
	"weather-monster/api/v1/forecasts"
	"weather-monster/api/v1/temperatures"
	"weather-monster/api/v1/webhooks"
	"weather-monster/config"
	"weather-monster/store"

	"github.com/go-chi/chi"
)

// Store holds new store connection
var Store *store.Conn

// Routes registered routes
func Routes(r chi.Router) {
	r.Method(http.MethodGet, "/", api.Handler(api.IndexHandeler))
	r.Get("/top", api.HealthHandeler)
	r.Route("/v1", Init)
}

// Init initializes all the v1 routes
func Init(r chi.Router) {
	r.Route("/cities", cities.Init)
	r.Route("/temperatures", temperatures.Init)
	r.Route("/forecasts", forecasts.Init)
	r.Route("/webhooks", webhooks.Init)

	fmt.Println(">> ", config.DBDriver, config.DBDataSource)
	Store = store.NewStore()
	fmt.Println("<< ", config.DBDriver, config.DBDataSource)
}
