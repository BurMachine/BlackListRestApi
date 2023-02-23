package handlers

import (
	"blacklistApi/internal/models"
	token2 "blacklistApi/pkg/token"
	"context"
	"encoding/json"
	"errors"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

func (h *Handlers) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			h.Logger.WithLevel(zerolog.WarnLevel).Msg("empty token")
			http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return
		}

		ok := h.storage.CheckToken(token)
		if !ok {
			h.Logger.WithLevel(zerolog.WarnLevel).Msg("token does not match")
			w.WriteHeader(http.StatusForbidden)
			http.Error(w, errors.New("invalid token").Error(), http.StatusForbidden)
		}
		h.Logger.Info().Msg("authentication passed")
		next.ServeHTTP(w, r.WithContext(context.Background()))
	}
}

// @Summary Auth
// @Tags auth
// @Description generate token
// @ID create-account
// @Accept json
// @Produce json
// @Param name body models.Name true "token generating by name"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth [post]
func (h *Handlers) Auth(w http.ResponseWriter, r *http.Request) {
	var body models.Name
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("parse body error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		h.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("unmarshalling error in parsing post request's body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := token2.GenerateToken(body.Name)
	if err != nil {
		h.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("auth invalid name")
		http.Error(w, errors.New("token gen error").Error(), http.StatusBadRequest)
		return
	}

	err = h.storage.AddToken(token)
	if err != nil {
		h.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("token adding error")
		http.Error(w, errors.New("token proc error").Error(), http.StatusInternalServerError)
		return
	}
	h.Logger.Info().Msg("token added")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
