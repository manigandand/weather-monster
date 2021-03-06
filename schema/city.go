package schema

import (
	err "errors"
	"strings"
	"weather-monster/pkg/errors"
)

// City it holds the properties of city schema
type City struct {
	BaseSchema
	Name      string  `json:"name" gorm:"unique_index" sql:"not null"`
	Slug      string  `json:"-"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Deleted   bool    `json:"deleted" sql:"default:false"`
}

// Ok implements the Ok interface, it validates city input
func (c *City) Ok() error {
	switch {
	case strings.TrimSpace(c.Name) == "":
		return errors.IsRequiredErr("name")
	case c.Deleted:
		return err.New("invalid request, you can't update deleted field")
	}

	return nil
}

// CityReq request payload to create/register a new city
type CityReq struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Ok implements the Ok interface, it validates city input
func (c *CityReq) Ok() error {
	switch {
	case strings.TrimSpace(c.Name) == "":
		return errors.IsRequiredErr("name")
	}
	return nil
}
