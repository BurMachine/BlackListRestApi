package handlers

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"net/http"
)

// @Summary Remove user from blacklist
// @Security ApiKeyAuth
// @Tags Removal
// @Description Remove user from blacklist
// @ID delete-from-blacklist
// @Accept  json
// @Produce  json
// @Param id query int true "ID of user to remove from blacklist"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /delete [get]
func (h *Handlers) Removal(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		h.Logger.WithLevel(zerolog.WarnLevel).Msg("empty removal_id")
		http.Error(w, errors.New("empty id").Error(), http.StatusBadRequest)
		return
	}
	err := h.storage.Remove(id)
	if err != nil {
		h.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("deletion error")
		http.Error(w, errors.New("not found").Error(), http.StatusNotFound)
		return
	}

	h.Logger.Info().Msg(fmt.Sprintf("[id=%v] deleted", id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
