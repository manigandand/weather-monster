package middleware

import (
	"context"
	"net/http"
	"strconv"
	"weather-monster/pkg/errors"
	"weather-monster/pkg/respond"
	"weather-monster/store"

	"github.com/go-chi/chi"
)

// Store holds new store connection
var Store *store.Conn

// Init ...
func Init(st *store.Conn) {
	Store = st
}

// CityRequired validates
func CityRequired(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cityIDStr := chi.URLParam(r, "cityID")
		cityID, er := strconv.Atoi(cityIDStr)
		if er != nil {
			respond.Fail(w, errors.BadRequest("invalid id").AddDebug(er))
			return
		}

		city, err := Store.City().GetByID(uint(cityID))
		if err != nil {
			respond.Fail(w, err)
			return
		}

		ctx := ContextWrapAll(r.Context(), map[interface{}]interface{}{
			"cityID": uint(cityID),
			"city":   city,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// ContextWrapAll is used to set the following values in the
// passed context
func ContextWrapAll(ctx context.Context, x map[interface{}]interface{}) context.Context {
	for key, value := range x {
		ctx = context.WithValue(ctx, key, value)
	}

	return ctx
}
