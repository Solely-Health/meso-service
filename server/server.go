package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/meso-org/meso/workers"
)

type Server struct {
	// TODO: make this an array that itterates through different service references and pushes them to a service store
	// we'll call this service registration or something
	WorkersSVC workers.Service
	router     chi.Router
}

// New - instantiates a new http server w/ router appended to it.
func New(ws workers.Service) *Server {
	s := &Server{
		WorkersSVC: ws,
	}

	r := chi.NewRouter()

	r.Use(accessControl)

	// Register worker module related endpoints
	r.Route("/worker", func(r chi.Router) {
		h := workerHandler{s.WorkersSVC}
		r.Mount("/v1", h.router())
	})

	s.router = r

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	// case repository.ErrUnknownWorker:
	// 	w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
