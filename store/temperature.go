package store

import (
	"fmt"
	"weather-monster/pkg/errors"
	"weather-monster/schema"
)

// TemperatureStore implements the cities interface
type TemperatureStore struct {
	*Conn
}

// NewTemperatureStore ...
func NewTemperatureStore(st *Conn) *TemperatureStore {
	ts := &TemperatureStore{st}
	go ts.createTableIfNotExists()
	return ts
}

func (ts *TemperatureStore) createTableIfNotExists() {
	if !ts.DB.HasTable(&schema.Temperature{}) {
		if err := ts.DB.CreateTable(&schema.Temperature{}).Error; err != nil {
			fmt.Println(err)
		}
	}
	if err := ts.DB.Model(&schema.Temperature{}).AddForeignKey("city_id", "cities(id)", "CASCADE", "RESTRICT").Error; err != nil {
		fmt.Println(err)
	}

	go ts.createIndexesIfNotExists()
}

func (ts *TemperatureStore) createIndexesIfNotExists() {
	scope := ts.DB.NewScope(&schema.Temperature{})
	commonIndexes := getCommonIndexes(scope.TableName())
	for k, v := range commonIndexes {
		if !scope.Dialect().HasIndex(scope.TableName(), k) {
			err := ts.DB.Model(&schema.Temperature{}).AddIndex(k, v).Error
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// All returns all the available temperature data for all the available cities
func (ts *TemperatureStore) All() ([]*schema.Temperature, *errors.AppError) {
	var tems []*schema.Temperature
	if err := ts.DB.Preload("City").Find(&tems).Error; err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return tems, nil
}

// Create a new temperature record for the city
func (ts *TemperatureStore) Create(tem *schema.Temperature) (*schema.Temperature, *errors.AppError) {
	if err := ts.DB.Save(tem).Error; err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return tem, nil
}

// GetByCityID returns all the temp data for the city
func (ts *TemperatureStore) GetByCityID(cityID uint) ([]*schema.Temperature, *errors.AppError) {
	var tems []*schema.Temperature
	if err := ts.DB.Preload("City").Find(&tems, "city_id=?", cityID).Error; err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return tems, nil
}
