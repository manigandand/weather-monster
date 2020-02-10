package schema

// Temperature schema holds the all the weather records for the city
type Temperature struct {
	BaseSchema
	// CityID is the foreign key; Belongs to `city` table.
	CityID    uint  `json:"city_id" sql:"not null"`
	Min       int   `json:"min"`
	Max       int   `json:"max"`
	Timestamp int64 `json:"timestamp"`
}

// Ok implements the Ok interface, it validates city input
func (c *Temperature) Ok() error {
	return nil
}
