package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/meso-org/meso/config"
	"github.com/meso-org/meso/positions"
	"github.com/meso-org/meso/repository"
	"gopkg.in/square/go-jose.v2/json"
)

type positionHandler struct {
	s positions.Service
}

func (h *positionHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/position", func(chi.Router) {
		r.Post("/", h.createPosition)
		r.Get("/", h.listPositions)
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
		Domain: "position",
		Ping:   "pong",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		encodeError(ctx, err, w)
		return
	}
}

func (h *positionHandler) createPosition(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request struct {
		FacilityID    string `json:"facilityID"`
		StartDateTime string `json:"startDateTime"`
		EndDateTime   string `json:"endDateTime"`
		Description   string `json:"description"`
		Title         string `json:"title"`
	}

	var response struct {
		Created    bool                  `json:"created"`
		PositionID repository.PositionID `json:"positionID"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("unable to decode json: %v", err)
	}

	// TODO: Double check whether or not the facility ID is real or not
	facilityID := repository.FacilityID(request.FacilityID)
	st, err := time.Parse(config.Dateformat, request.StartDateTime)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	et, err := time.Parse(config.Dateformat, request.EndDateTime)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	startDateTime := repository.JSONTime(st)
	endDateTime := repository.JSONTime(et)

	response.PositionID, err = h.s.CreateNewPosition(facilityID, startDateTime, endDateTime, request.Description, request.Title)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}
	response.Created = true

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *positionHandler) listPositions(w http.ResponseWriter, r *http.Request) {
	var err error
	var response struct {
		Positions []*repository.Position `json:"positions"`
	}
	response.Positions, err = h.s.FindAllPositions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
