package store

// ForecastStore implements the cities interface
type ForecastStore struct {
	*Conn
}

// NewForecastStore ...
func NewForecastStore(st *Conn) *ForecastStore {
	return &ForecastStore{st}
}
