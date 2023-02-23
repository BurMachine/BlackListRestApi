package handlers

import (
	"blacklistApi/internal/database"
	gorilla "github.com/gorilla/mux"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"

	_ "blacklistApi/docs"
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
	mux.HandleFunc("/auth", h.Auth).Methods(http.MethodPost)
	mux.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Generate Swagger documentation
	swagHandler := httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/swagger.json"), // The url pointing to API definition"
	)

	mux.Handle("/swagger/doc.json", swagHandler)

}

// response

type errorResponse struct {
	Message string `json:"message"`
}
