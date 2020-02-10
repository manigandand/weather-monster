package schema

// Webhook schema holds the all the webhook records,
// to post the weather reports
type Webhook struct {
	baseSchema
	// CityID is the foreign key; Belongs to `city` table.
	CityID      uint   `json:"city_id" sql:"not null"`
	CallbackURL string `json:"callback_url" sql:"not null"`
	Deleted     bool   `json:"deleted" sql:"default:false"`
}

// Ok implements the Ok interface, it validates city input
func (c *Webhook) Ok() error {
	return nil
}
