package handlers

import (
	"blacklistApi/internal/database"
	gorilla "github.com/gorilla/mux"
	"net/http"
)

type Handlers struct {
	storage *database.Storage
}

func New(storage *database.Storage) *Handlers {
	return &Handlers{storage: storage}
}

func (h *Handlers) RegisteringHandlers(mux *gorilla.Router) {
	mux.HandleFunc("/add", h.Addition).Methods(http.MethodPost)
	mux.HandleFunc("/delete", h.Removal).Methods(http.MethodGet)
	mux.HandleFunc("/search", h.Search).Methods(http.MethodGet)
}
