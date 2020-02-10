package store

import (
	"fmt"
	"weather-monster/schema"
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
}
