package main

import (
	"fmt"
	"net/http"
	"weather-monster/api"
	v1 "weather-monster/api/v1"
	"weather-monster/config"
	appmiddleware "weather-monster/middleware"
	"weather-monster/pkg/trace"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

var (
	name    = "weather_monster"
	version = "1.0.0"
)

func main() {
	api.InitService(name, version)
	trace.Setup(config.Env)

	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{
			"Origin", "Authorization", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Header", "Accept",
			"Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders: []string{
			"Content-Length", "Access-Control-Allow-Origin", "Origin",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// cross & loger middleware
	router.Use(cors.Handler)
	router.Use(
		middleware.Logger,
		appmiddleware.Recoverer,
	)

	// Initialize the version 1 routes of the API
	router.Get("/", api.IndexHandeler)
	router.Get("/top", api.HealthHandeler)
	router.Route("/v1", v1.Init)

	trace.Log.Infof("Starting %s:%s on port :%s\n", name, version, config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router)
}
