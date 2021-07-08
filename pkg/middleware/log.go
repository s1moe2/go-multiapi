package middleware

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := statusRecorder{w, http.StatusOK}
		next.ServeHTTP(&rec, r)
		log.WithFields(log.Fields{
			"req":    fmt.Sprintf("%s %s", r.Method, r.RequestURI),
			"status": rec.status,
		}).Info("handled request")
	})
}
