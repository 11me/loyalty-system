package middleware

import (
	"loyalty-system/pkg/logger"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		responseRecorder := &statusRecorder{
			ResponseWriter: w,
			status:         200,
		}
		log := logger.GetLogger()
		ctx := r.Context()
		timeStart := time.Now()
		next.ServeHTTP(responseRecorder, r)
		timeEnd := time.Since(timeStart).Milliseconds()

		l := log.WithFields(logger.Fields{
			"duration_ms": timeEnd,
			"request_id":  ctx.Value("Request-ID"),
			"status_code": responseRecorder.status,
			"method":      r.Method,
			"user-agent":  r.UserAgent(),
			"path":        r.URL.Path,
		})

		if responseRecorder.status >= 400 && responseRecorder.status < 500 {
			l.Warn("Incoming request can not be processed")
		} else if responseRecorder.status >= 500 {
			l.Error("Incoming request failed")
		} else {
			l.Info("Incoming request succeeded")
		}
	}
	return http.HandlerFunc(fn)
}
