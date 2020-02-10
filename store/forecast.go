package store

import (
	"time"
	"weather-monster/pkg/errors"
	"weather-monster/schema"
)

// ForecastStore implements the cities interface
type ForecastStore struct {
	*Conn
}

// NewForecastStore ...
func NewForecastStore(st *Conn) *ForecastStore {
	return &ForecastStore{st}
}

// ByCityID returns the forcast result for the given city
func (fs *ForecastStore) ByCityID(cityID uint) (*schema.Forecast, *errors.AppError) {
	var (
		min, max float64
		sample   int
	)

	query := `
	COALESCE(ROUND(AVG(min),2),0) as min,
	COALESCE(ROUND(AVG(max),2),0) as max,
	COUNT(*) as sample
	`
	now := time.Now().Unix()
	if err := fs.DB.Model(&schema.Temperature{}).Select(query).
		Where("city_id=? and timestamp  BETWEEN ? AND ?", cityID, now-86400, now).Row().
		Scan(&min, &max, &sample); err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return &schema.Forecast{
		CityID: cityID,
		Min:    min,
		Max:    max,
		Sample: sample,
	}, nil
}
