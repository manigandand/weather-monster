package cities

import (
	"net/http"
	"weather-monster/api"
	"weather-monster/pkg/errors"
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
	r.Route("/{cityID}", func(r chi.Router) {
		// ROUTE: {host}/v1/cities/:cityID
		// r.Use(ArticleCtx)
		r.Method(http.MethodGet, "/", api.Handler(auctionHandler))
		r.Method(http.MethodPatch, "/", api.Handler(auctionHandler))
		r.Method(http.MethodDelete, "/", api.Handler(auctionHandler))
	})
}

func auctionHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	return nil
}
