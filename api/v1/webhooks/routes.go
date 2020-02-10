package webhooks

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

	// ROUTE: {host}/v1/webhooks
	r.Method(http.MethodGet, "/", api.Handler(getAllWebhooksHandler))
	r.Method(http.MethodPost, "/", api.Handler(saveCityWebhooksHandler))
	r.With(middleware.WebhookRequired).
		Route("/{webhookID:[0-9]+}", webhookIDSubRoutes)
}

// ROUTE: {host}/v1/webhooks/:webhookID/*
func webhookIDSubRoutes(r chi.Router) {
	r.Method(http.MethodGet, "/", api.Handler(getWebhookHandler))
	r.Method(http.MethodDelete, "/", api.Handler(deleteWebhookHandler))
}
