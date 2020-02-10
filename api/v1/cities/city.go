package cities

import (
	"net/http"
	"weather-monster/pkg/errors"
	"weather-monster/pkg/respond"
	"weather-monster/schema"
	"weather-monster/utils"
)

// returns all the available cities
func getAllCitiesHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	cities, err := store.City().All()
	if err != nil {
		return err
	}

	respond.OK(w, cities)
	return nil
}

// creates a new city
func createCityHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	var input schema.CityReq

	if err := utils.Decode(r, &input); err != nil {
		return errors.BadRequest(err.Error()).AddDebug(err)
	}

	city, err := store.City().Create(&input)
	if err != nil {
		return err
	}

	respond.Created(w, city)
	return nil
}
