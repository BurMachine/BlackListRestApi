package handlers

import (
	"blacklistApi/internal/database"
	gorilla "github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
)

type Handlers struct {
	storage *database.Storage
	Logger  *zerolog.Logger
}

func New(storage *database.Storage) *Handlers {
	return &Handlers{storage: storage}
}

func (h *Handlers) RegisteringHandlers(mux *gorilla.Router) {
	mux.HandleFunc("/add", h.AuthMiddleware(h.Addition)).Methods(http.MethodPost)
	mux.HandleFunc("/delete", h.AuthMiddleware(h.Removal)).Methods(http.MethodGet)
	mux.HandleFunc("/search", h.AuthMiddleware(h.Search)).Methods(http.MethodGet)
	mux.HandleFunc("/auth", h.Auth).Methods(http.MethodGet)
}

// response

type errorResponse struct {
	Message string `json:"message"`
}
