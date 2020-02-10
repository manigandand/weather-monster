package forecasts

import (
	"net/http"
	"weather-monster/api"
	"weather-monster/middleware"
	"weather-monster/pkg/errors"
	"weather-monster/pkg/respond"
	"weather-monster/schema"
	appstore "weather-monster/store"

	"github.com/go-chi/chi"
)

// store holds shared store conn from the api
var store *appstore.Conn

// Init initializes all the v1 routes
func Init(r chi.Router) {
	store = api.Store
	// ROUTE: {host}/v1/forecasts/:cityID
	r.With(middleware.CityRequired).
		Method(http.MethodGet, "/{cityID:[0-9]+}", api.Handler(WeatherForecastHandler))
}

// WeatherForecastHandler returns the avg weather report for last 24 hrs for the city
func WeatherForecastHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()
	city, _ := ctx.Value("city").(*schema.City)
	forecast, err := store.Forecast().ByCityID(city.ID)
	if err != nil {
		return err
	}
	forecast.City = city

	respond.OK(w, forecast)
	return nil
}
