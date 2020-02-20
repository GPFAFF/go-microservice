package home

import (
	"log"
	"net/http"
	"time"
)

const message = "Microservice Live"

type handlers struct {
	logger *log.Logger
}

//Page builds the homepage route.
func (h *handlers) Page(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

//Logger creates the middleware.
func (h *handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.Logger.Printf("Request was processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
