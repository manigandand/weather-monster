package schema

import (
	"net/url"
	"strings"
	"weather-monster/pkg/errors"
)

// Webhook schema holds the all the webhook records,
// to post the weather reports
type Webhook struct {
	BaseSchema
	// CityID is the foreign key; Belongs to `city` table.
	CityID      uint   `json:"city_id" sql:"not null"`
	City        *City  `json:"city,omitempty"`
	CallbackURL string `json:"callback_url" sql:"not null"`
	Deleted     bool   `json:"deleted" sql:"default:false"`
}

// Ok implements the Ok interface, it validates city input
func (c *Webhook) Ok() error {
	switch {
	case c.CityID == 0:
		return errors.IsRequiredErr("city id")
	case strings.TrimSpace(c.CallbackURL) == "":
		return errors.IsRequiredErr("call back url")
	}
	if _, err := url.Parse(c.CallbackURL); err != nil {
		return err
	}

	return nil
}
