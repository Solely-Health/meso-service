package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/meso-org/meso/repository"
	"github.com/meso-org/meso/workers"
)

type workerHandler struct {
	s workers.Service
}

func (h *workerHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/worker", func(chi.Router) {
		// TODO: this is a dummy login handler for later
		r.Post("/login", h.loginWorker)
		r.Post("/", h.registerWorker)
		r.Get("/", h.listWorkers)
		r.Route("/{workerID}", func(r chi.Router) {
			r.Get("/", h.findWorker)
		})
		/*
			if we were to add more sub routing:
			r.Route("/pattern", func(chi.Router) {
				r.Verb("/pattern", handlerFunc)
			})
		*/
	})
	r.Get("/ping", h.testPing)
	return r
}

func (h *workerHandler) testPing(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var response = struct {
		Domain string `json:"domain"`
		Ping   string `json:"ping"`
	}{
		Domain: "worker",
		Ping:   "pong",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		encodeError(ctx, err, w)
		return
	}
}

func (h *workerHandler) loginWorker(w http.ResponseWriter, r *http.Request) {
	var err error

	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var response struct {
		User  *repository.Worker `json:"worker"`
		Token string             `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("unable to decode json: %v", err)
	}

	parsedEmail := repository.Email(request.Email)

	response.User, err = h.s.FindWorkerByEmail(parsedEmail)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if response.User.Email == parsedEmail {
		response.Token = "FAKEOAUTHTOKEN1234567890"
		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (h *workerHandler) registerWorker(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email      string `json:"email"`
		FirstName  string `json:"firstName"`
		LastName   string `json:"lastName"`
		Occupation string `json:"occupation"`
		License    string `json:"license"`
	}

	var response struct {
		ID repository.WorkerID `json:"workerId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("unable to decode json: %v", err)
	}

	fmt.Println("heres the request string ", request.Email, request.FirstName, request.LastName, request.Occupation, request.License)

	id, err := h.s.RegisterNewWorker(request.Email, request.FirstName, request.LastName, request.Occupation, request.License)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	response.ID = id

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *workerHandler) findWorker(w http.ResponseWriter, r *http.Request) {
	var err error

	var response struct {
		Worker *repository.Worker `json:"worker"`
	}

	workerID := repository.WorkerID(chi.URLParam(r, "workerID"))

	fmt.Println("heres the request string ", workerID)

	response.Worker, err = h.s.FindWorkerByID(workerID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *workerHandler) listWorkers(w http.ResponseWriter, r *http.Request) {
	var err error
	var response struct {
		Workers []*repository.Worker `json:"workers"`
	}
	response.Workers, err = h.s.FindAllWorkers()
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
