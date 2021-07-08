package http

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type internalError struct {
	Message string
}

type AppError struct {
	Errors []string
}

// RespondJSON is an helper that takes care of the
// HTTP response part of a request handler
func RespondJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Errorln("respondJSON", err)
	}
}

// RespondInternalError is an helper similar to respondError but responds
// with a default internal error code and payload
func RespondInternalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(w).Encode(&internalError{
		Message: "Oops! Something went wrong on our side.",
	})
	if err != nil {
		log.Errorln("respondInternalError", err)
	}
}

func RespondNoContent(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}
