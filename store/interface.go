package store

import (
	"weather-monster/pkg/errors"
	"weather-monster/schema"
)

// Store global store interface - provides db intercae methods
// for diff entities
type Store interface {
	City() Cities
	Temperature() Temperatures
	Forecast() Forecasts
	Webhook() Webhooks
}

// Cities store interface expose the city db methods
type Cities interface {
	All() ([]*schema.City, *errors.AppError)
	Create(req *schema.CityReq) (*schema.City, *errors.AppError)
}

// Temperatures store interface expose the Temperatures db methods
type Temperatures interface{}

// Forecasts store interface expose the Forecasts db methods
type Forecasts interface{}

// Webhooks store interface expose the Webhooks db methods
type Webhooks interface{}
