package cities

import (
	"net/http"
	"weather-monster/api"
	"weather-monster/middleware"
	appstore "weather-monster/store"

	"github.com/go-chi/chi"
)

// store holds shared store conn from the api
var store *appstore.Conn

// Init initializes all the v1 routes
func Init(r chi.Router) {
	store = api.Store

	// ROUTE: {host}/v1/cities
	r.Method(http.MethodGet, "/", api.Handler(getAllCitiesHandler))
	r.Method(http.MethodPost, "/", api.Handler(createCityHandler))
	r.With(middleware.CityRequired).
		Route("/{cityID:[0-9]+}", cityIDSubRoutes)
}

// ROUTE: {host}/v1/cities/:cityID/*
func cityIDSubRoutes(r chi.Router) {
	r.Method(http.MethodGet, "/", api.Handler(getCityHandler))
	r.Method(http.MethodPatch, "/", api.Handler(updateCityHandler))
	r.Method(http.MethodDelete, "/", api.Handler(deleteCityHandler))

	r.Method(http.MethodGet, "/temperature", api.Handler(getCityTemperatureHandler))
}
