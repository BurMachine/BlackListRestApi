package handlers

import (
	"blacklistApi/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"net/http"
)

// @Summary Search users in blacklist
// @Security ApiKeyAuth
// @Tags Search
// @Description Search users in blacklist
// @ID search-users
// @Accept  json
// @Produce  json
// @Param number query string false "user phone"
// @Param name query string false "user name"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /search [get]
func (h *Handlers) Search(w http.ResponseWriter, r *http.Request) {
	number := r.URL.Query().Get("number")
	name := r.URL.Query().Get("name")
	if number == "" && name == "" {
		h.Logger.WithLevel(zerolog.WarnLevel).Msg("no params to search")
		http.Error(w, errors.New("empty number").Error(), http.StatusBadRequest)
		return
	}
	var users []models.Addiction
	var err error
	if number != "" {
		users, err = h.storage.Search("phone_number", number)
		if err != nil {
			h.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("search error")
			http.Error(w, errors.New("search error").Error(), http.StatusNotFound)
			return
		}
		h.Logger.Info().Msg(fmt.Sprintf("user(s) with phone - [%v] found", number))
	} else {
		users, err = h.storage.Search("user_name", name)
		if err != nil {
			h.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("search error")
			http.Error(w, errors.New("search error").Error(), http.StatusNotFound)
			return
		}
		h.Logger.Info().Msg(fmt.Sprintf("user(s) with name - [%v] found", name))
	}
	resp, err := json.Marshal(users)
	if err != nil {
		h.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("search result marshalling error")
		http.Error(w, errors.New("search error").Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
