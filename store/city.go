package store

import (
	"fmt"
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
