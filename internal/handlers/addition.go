package handlers

import (
	"blacklistApi/internal/models"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"time"
)

// @Summary Add user to blacklist
// @Security ApiKeyAuth
// @Tags Addition
// @Description Add user to blacklist
// @ID add-into-blacklist
// @Accept  json
// @Produce  json
// @Param input body models.Addiction true "user info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /add [post]
func (h *Handlers) Addition(w http.ResponseWriter, r *http.Request) {
	var body models.Addiction
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
	body.Date = time.Now()
	err = h.storage.Add(body)
	if err != nil {
		h.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("sql adding error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Logger.Info().Msg(fmt.Sprintf("[%v] added", body))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
