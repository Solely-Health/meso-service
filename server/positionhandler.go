package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/meso-org/meso/positions"
	"gopkg.in/square/go-jose.v2/json"
)

type positionHandler struct {
	s positions.Service
}

func (h *positionHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/position", func(chi.Router) {
		r.Post("/", h.createPosition)
		r.Get("/ping", h.testPing)
	})

	return r
}

func (h *positionHandler) testPing(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var response = struct {
		Domain string `json:"domain"`
		Ping   string `json:"ping"`
	}{
		Domain: "facility",
		Ping:   "pong",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		encodeError(ctx, err, w)
		return
	}
}

func (h *positionHandler) createPosition(w http.ResponseWriter, r *http.Request) {
}
