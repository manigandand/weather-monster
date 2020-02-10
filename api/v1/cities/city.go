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

func getCityHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()
	city, _ := ctx.Value("city").(*schema.City)

	respond.OK(w, city)
	return nil
}

func updateCityHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	var input schema.City
	ctx := r.Context()
	city, _ := ctx.Value("city").(*schema.City)

	if err := utils.Decode(r, &input); err != nil {
		return errors.BadRequest(err.Error()).AddDebug(err)
	}

	updated, err := store.City().Update(city, &input)
	if err != nil {
		return err
	}

	respond.OK(w, updated)
	return nil
}

func deleteCityHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()
	cityID, _ := ctx.Value("cityID").(uint)

	if err := store.City().Delete(cityID); err != nil {
		return err
	}
	respond.NoContent(w)
	return nil
}

func getCityTemperatureHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()
	cityID, _ := ctx.Value("cityID").(uint)

	temps, err := store.Temperature().GetByCityID(cityID)
	if err != nil {
		return err
	}

	respond.OK(w, temps)
	return nil
}

func getCityWebhookHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()
	cityID, _ := ctx.Value("cityID").(uint)

	webhooks, err := store.Webhook().GetByCityID(cityID)
	if err != nil {
		return err
	}

	respond.OK(w, webhooks)
	return nil
}
