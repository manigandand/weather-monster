package temperatures

import (
	"net/http"
	"time"
	"weather-monster/pkg/errors"
	"weather-monster/pkg/respond"
	"weather-monster/schema"
	"weather-monster/utils"
)

// returns all the available temperature data for all the available cities
func getAllTemperaturesHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	cities, err := store.Temperature().All()
	if err != nil {
		return err
	}

	respond.OK(w, cities)
	return nil
}

// creates a new temperature record for the city
func saveCityTemperaturesHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	var input schema.Temperature
	if err := utils.Decode(r, &input); err != nil {
		return errors.BadRequest(err.Error()).AddDebug(err)
	}
	if _, err := store.City().GetByID(input.CityID); err != nil {
		return err
	}

	input.Timestamp = time.Now().Unix()
	temp, err := store.Temperature().Create(&input)
	if err != nil {
		return err
	}

	respond.Created(w, temp)
	return nil
}
