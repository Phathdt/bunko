package web

import (
	"bunko"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func NewHandler(store bunko.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}
	h.Route("/threads", func(r chi.Router) {
		r.Get("/", h.ThreadList())
	})
	return h
}

type Handler struct {
	*chi.Mux
	store bunko.Store
}

func (h *Handler) ThreadList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		threads, _ := h.store.Threads()

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(threads)
	}
}
