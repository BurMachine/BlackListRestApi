package handlers

import (
	"blacklistApi/internal/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (h *Handlers) Addition(w http.ResponseWriter, r *http.Request) {
	var body models.Addiction
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("parse body error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		log.Println("unmarshalling error in parsing post request's body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.storage.Add(body)
	if err != nil {
		log.Println("insert error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("added"))
}
