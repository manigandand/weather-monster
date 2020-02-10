package store

import (
	"fmt"
	"weather-monster/pkg/errors"
	"weather-monster/schema"

	"github.com/jinzhu/gorm"
)

// WebhookStore implements the cities interface
type WebhookStore struct {
	*Conn
}

// NewWebhookStore ...
func NewWebhookStore(st *Conn) *WebhookStore {
	ws := &WebhookStore{st}
	go ws.createTableIfNotExists()
	return ws
}

func (ws *WebhookStore) createTableIfNotExists() {
	if !ws.DB.HasTable(&schema.Webhook{}) {
		if err := ws.DB.CreateTable(&schema.Webhook{}).Error; err != nil {
			fmt.Println(err)
		}
	}
	if err := ws.DB.Model(&schema.Webhook{}).AddForeignKey("city_id", "cities(id)", "CASCADE", "RESTRICT").Error; err != nil {
		fmt.Println(err)
	}

	go ws.createIndexesIfNotExists()
}

func (ws *WebhookStore) createIndexesIfNotExists() {
	scope := ws.DB.NewScope(&schema.Webhook{})
	commonIndexes := getCommonIndexes(scope.TableName())
	for k, v := range commonIndexes {
		if !scope.Dialect().HasIndex(scope.TableName(), k) {
			err := ws.DB.Model(&schema.Webhook{}).AddIndex(k, v).Error
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	uniqueIndexes := map[string][]string{
		"idx_webhooks_city_id_callback_url": []string{"city_id", "callback_url", "deleted_at"},
	}
	for k, v := range uniqueIndexes {
		if !scope.Dialect().HasIndex(scope.TableName(), k) {
			if err := ws.DB.Model(&schema.Webhook{}).AddUniqueIndex(k, v...).Error; err != nil {
				fmt.Println(err)
			}
		}
	}
}

// All returns all the webhooks
func (ws *WebhookStore) All() ([]*schema.Webhook, *errors.AppError) {
	var webhooks []*schema.Webhook
	if err := ws.DB.Preload("City").Find(&webhooks, "deleted=?", false).Error; err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return webhooks, nil
}

// Create a new webhook
func (ws *WebhookStore) Create(req *schema.Webhook) (*schema.Webhook, *errors.AppError) {
	if recordExists("webhooks", fmt.Sprintf("city_id=%d and callback_url='%s' and deleted_at=null", req.CityID, req.CallbackURL)) {
		return nil, errors.BadRequest("you already suscribed for this city")
	}
	if err := ws.DB.Save(req).Error; err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return req, nil
}

// GetByID returns the matched record for the given id
func (ws *WebhookStore) GetByID(webhookID uint) (*schema.Webhook, *errors.AppError) {
	var webhook schema.Webhook
	if err := ws.DB.First(&webhook, "id=? and deleted=?", webhookID, false).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.BadRequest("invalid webhook id").AddDebug(err)
		}
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return &webhook, nil
}

// GetByCityID returns all the webhooks data for the city
func (ws *WebhookStore) GetByCityID(cityID uint) ([]*schema.Webhook, *errors.AppError) {
	var webhooks []*schema.Webhook
	if err := ws.DB.Preload("City").Find(&webhooks, "city_id=? and deleted=?", cityID, false).Error; err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return webhooks, nil
}

// Delete soft deletes the city for the given id
func (ws *WebhookStore) Delete(webhookID uint) *errors.AppError {
	if err := ws.DB.Delete(&schema.Webhook{}, "id=?", webhookID).Error; err != nil {
		return errors.InternalServerStd().AddDebug(err)
	}

	return nil
}
