package store

import (
	"fmt"
	"weather-monster/pkg/errors"
	"weather-monster/schema"
)

// CityStore implements the cities interface
type CityStore struct {
	*Conn
}

// NewCityStore ...
func NewCityStore(st *Conn) *CityStore {
	cs := &CityStore{st}
	go cs.createTableIfNotExists()
	return cs
}

func (cs *CityStore) createTableIfNotExists() {
	if !cs.DB.HasTable(&schema.City{}) {
		if err := cs.DB.CreateTable(&schema.City{}).Error; err != nil {
			fmt.Println(err)
		}
	}

	go cs.createIndexesIfNotExists()
}

func (cs *CityStore) createIndexesIfNotExists() {
	scope := cs.DB.NewScope(&schema.City{})
	commonIndexes := getCommonIndexes(scope.TableName())
	for k, v := range commonIndexes {
		if !scope.Dialect().HasIndex(scope.TableName(), k) {
			err := cs.DB.Model(&schema.City{}).AddIndex(k, v).Error
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	uniqueIndexes := map[string][]string{
		"idx_cities_name": []string{"name"},
	}
	for k, v := range uniqueIndexes {
		if !scope.Dialect().HasIndex(scope.TableName(), k) {
			if err := cs.DB.Model(&schema.City{}).AddUniqueIndex(k, v...).Error; err != nil {
				fmt.Println(err)
			}
		}
	}
}

// All returns all the cities
func (cs *CityStore) All() ([]*schema.City, *errors.AppError) {
	var cities []*schema.City
	if err := cs.DB.Find(&cities, "deleted=?", false).Error; err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return cities, nil
}

// Create a new city
func (cs *CityStore) Create(req *schema.CityReq) (*schema.City, *errors.AppError) {
	city := &schema.City{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
	if err := cs.DB.Save(city).Error; err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return city, nil
}
