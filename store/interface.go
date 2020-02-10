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
	GetByID(cityID uint) (*schema.City, *errors.AppError)
	Update(city *schema.City, update *schema.City) (*schema.City, *errors.AppError)
	Delete(cityID uint) *errors.AppError
}

// Temperatures store interface expose the Temperatures db methods
type Temperatures interface {
	All() ([]*schema.Temperature, *errors.AppError)
	Create(tem *schema.Temperature) (*schema.Temperature, *errors.AppError)
	GetByCityID(cityID uint) ([]*schema.Temperature, *errors.AppError)
}

// Forecasts store interface expose the Forecasts db methods
type Forecasts interface {
	ByCityID(cityID uint) (*schema.Forecast, *errors.AppError)
}

// Webhooks store interface expose the Webhooks db methods
type Webhooks interface {
	All() ([]*schema.Webhook, *errors.AppError)
	Create(req *schema.Webhook) (*schema.Webhook, *errors.AppError)
	GetByID(webhookID uint) (*schema.Webhook, *errors.AppError)
	GetByCityID(cityID uint) ([]*schema.Webhook, *errors.AppError)
	Delete(webhookID uint) *errors.AppError
}
