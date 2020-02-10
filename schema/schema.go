package schema

import "time"

// BaseSchema ...
type BaseSchema struct {
	ID        uint       `json:"id" sql:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" sql:"default:current_timestamp"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// Ok interface/method validates the struct data
type Ok interface {
	Ok() error
}
