package temperatures

import (
	"net/http"
	"weather-monster/api"
	appstore "weather-monster/store"

	"github.com/go-chi/chi"
)

// store holds shared store conn from the api
var store *appstore.Conn

// Init initializes all the v1 routes
func Init(r chi.Router) {
	store = api.Store

	// ROUTE: {host}/v1/temperatures
	r.Method(http.MethodGet, "/", api.Handler(getAllTemperaturesHandler))
	r.Method(http.MethodPost, "/", api.Handler(saveCityTemperaturesHandler))
}
