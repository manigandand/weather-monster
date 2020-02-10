package middleware

import (
	"net/http"
	"strconv"
	"weather-monster/pkg/errors"
	"weather-monster/pkg/respond"

	"github.com/go-chi/chi"
)

// WebhookRequired validates the resource
func WebhookRequired(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		webhookIDStr := chi.URLParam(r, "webhookID")
		webhookID, er := strconv.Atoi(webhookIDStr)
		if er != nil {
			respond.Fail(w, errors.BadRequest("invalid id").AddDebug(er))
			return
		}

		webhook, err := Store.Webhook().GetByID(uint(webhookID))
		if err != nil {
			respond.Fail(w, err)
			return
		}

		ctx := ContextWrapAll(r.Context(), map[interface{}]interface{}{
			"webhookID": uint(webhookID),
			"webhook":   webhook,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
