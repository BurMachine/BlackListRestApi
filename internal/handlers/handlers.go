package handlers

import (
	gorilla "github.com/gorilla/mux"
	"net/http"
)

type Handlers struct {
}

func New() *Handlers {
	return &Handlers{}
}

func (h *Handlers) RegisteringHandlers(mux *gorilla.Router) {
	mux.HandleFunc("/add", h.Addition).Methods(http.MethodPost)
	mux.HandleFunc("/delete", h.Removal).Methods(http.MethodGet)
	mux.HandleFunc("/search", h.Search).Methods(http.MethodGet)
}
