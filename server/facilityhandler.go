package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/meso-org/meso/facilities"
	"github.com/meso-org/meso/repository"
	"gopkg.in/square/go-jose.v2/json"
)

type facilityHandler struct {
	s facilities.Service
}

func (h *facilityHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/facility", func(chi.Router) {
		r.Post("/", h.registerFacility)
	})

	return r
}

func (h *facilityHandler) registerFacility(w http.ResponseWriter, r *http.Request) {
	var err error
	var request struct {
		Email        string `json:"email"`
		FacilityName string `json:"facilityName"`
	}

	var response struct {
		ID repository.FacilityID
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("unable to decode json: %v", err)
	}

	response.ID, err = h.s.RegisterNewFacility(request.FacilityName, request.Email)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
