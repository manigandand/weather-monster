package store

import (
	"fmt"
	"log"
	"weather-monster/config"
	"weather-monster/schema"

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
	db.AutoMigrate(
		&schema.City{},
		&schema.Temperature{},
		&schema.Webhook{},
	)
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
	conn := &Conn{
		DB: dbConn,
	}
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

func getCommonIndexes(tableName string) map[string]string {
	idx := fmt.Sprintf("idx_%s", tableName)
	return map[string]string{
		fmt.Sprintf("%s_created_at", idx): "created_at",
		fmt.Sprintf("%s_updated_at", idx): "updated_at",
		fmt.Sprintf("%s_deleted_at", idx): "deleted_at",
	}
}

// recordExists should check if record is avail or not for particular table
// based on the given condition.
func recordExists(tableName, where string) (exists bool) {
	baseQ := fmt.Sprintf("select 1 from %s where %v", tableName, where)
	dbConn.Raw(fmt.Sprintf("select exists (%v)", baseQ)).Row().Scan(&exists)
	return
}
