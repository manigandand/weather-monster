package temperatures

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"weather-monster/pkg/errors"
	"weather-monster/pkg/respond"
	"weather-monster/pkg/trace"

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

	city, err := store.City().GetByID(input.CityID)
	if err != nil {
		return err
	}

	input.Timestamp = time.Now().Unix()
	temp, err := store.Temperature().Create(&input)
	if err != nil {
		return err
	}
	go postToWebhooks(city, temp)

	respond.Created(w, temp)
	return nil
}

func postToWebhooks(city *schema.City, temp *schema.Temperature) {
	webhooks, err := store.Webhook().GetByCityID(temp.CityID)
	if err != nil {
		return
	}

	data := map[string]interface{}{
		"city_id": temp.CityID,
		"city": map[string]interface{}{
			"name":      city.Name,
			"latitude":  city.Latitude,
			"longitude": city.Longitude,
		},
		"min":       temp.Min,
		"max":       temp.Max,
		"timestamp": temp.Timestamp,
	}

	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(data); err != nil {
		trace.Log.Info(err)
		return
	}

	for _, hook := range webhooks {
		go func(url string, body *bytes.Buffer) {
			trace.Log.Info("posting to > ", url)
			req, err := http.NewRequest(http.MethodPost, url, body)
			if err != nil {
				return
			}
			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return
			}

			defer resp.Body.Close()
			trace.Log.Info(url, " > ", resp.StatusCode)
		}(hook.CallbackURL, body)
	}
}
