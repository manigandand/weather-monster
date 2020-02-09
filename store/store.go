package store

import (
	"log"
	"weather-monster/config"

	"github.com/jinzhu/gorm"
	// gorm postgres connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbConn *gorm.DB

func init() {
	db, err := gorm.Open(config.DBDriver, config.DBDataSource)
	if err != nil {
		log.Fatal(err)
	}
	dbConn = db
}

// Conn struct holds the store connection
type Conn struct {
	DB              *gorm.DB
	CityConn        Cities
	TemperatureConn Temperatures
	ForecastsConn   Forecasts
	WebhooksConn    Webhooks
}

// NewStore inits new store connection
func NewStore() *Conn {
	conn := new(Conn)
	conn.CityConn = NewCityStore(conn)
	conn.TemperatureConn = NewTemperatureStore(conn)
	conn.ForecastsConn = NewForecastStore(conn)
	conn.WebhooksConn = NewWebhookStore(conn)

	return conn
}

// City implements the store interface and it returns the Cities interface
func (s *Conn) City() Cities {
	return s.CityConn
}

// Temperature implements the store interface and it returns the Temperatures interface
func (s *Conn) Temperature() Temperatures {
	return s.TemperatureConn
}

// Forecast implements the store interface and it returns the Forecasts interface
func (s *Conn) Forecast() Forecasts {
	return s.ForecastsConn
}

// Webhook implements the store interface and it returns the Webhooks interface
func (s *Conn) Webhook() Webhooks {
	return s.WebhooksConn
}
