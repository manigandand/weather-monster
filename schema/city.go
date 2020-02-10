package schema

// City it holds the properties of city schema
type City struct {
	baseSchema
	Name      string  `json:"name" gorm:"unique_index" sql:"not null"`
	Slug      string  `json:"slug"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Deleted   bool    `json:"deleted" sql:"default:false"`
}

// Ok implements the Ok interface, it validates city input
func (c *City) Ok() error {
	return nil
}
