package webhooks

import (
	"net/http"
	"weather-monster/pkg/errors"
	"weather-monster/pkg/respond"
	"weather-monster/schema"
	"weather-monster/utils"
)

// returns all the registered webhooks
func getAllWebhooksHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	webhooks, err := store.Webhook().All()
	if err != nil {
		return err
	}

	respond.OK(w, webhooks)
	return nil
}

// registeres a new webhook record for the city
func saveCityWebhooksHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	var input schema.Webhook
	if err := utils.Decode(r, &input); err != nil {
		return errors.BadRequest(err.Error()).AddDebug(err)
	}
	if _, err := store.City().GetByID(input.CityID); err != nil {
		return err
	}

	webhook, err := store.Webhook().Create(&input)
	if err != nil {
		return err
	}

	respond.Created(w, webhook)
	return nil
}

func getWebhookHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()
	webhook, _ := ctx.Value("webhook").(*schema.Webhook)

	respond.OK(w, webhook)
	return nil
}

func deleteWebhookHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()
	webhookID, _ := ctx.Value("webhookID").(uint)

	if err := store.Webhook().Delete(webhookID); err != nil {
		return err
	}
	respond.NoContent(w)
	return nil
}
