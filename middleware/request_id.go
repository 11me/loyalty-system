package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		ctx := r.Context()
		if reqID == "" {
			reqID = uuid.New().String()
			ctx = context.WithValue(ctx, "Request-ID", reqID)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
