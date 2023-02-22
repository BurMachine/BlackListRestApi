package handlers

import (
	"errors"
	"log"
	"net/http"
)

func (h *Handlers) Removal(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("Invalid id received(empty)")
		http.Error(w, errors.New("empty id").Error(), http.StatusBadRequest)
		return
	}
	err := h.storage.Remove(id)
	if err != nil {
		log.Println("error remove")
		http.Error(w, errors.New("removal error").Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("removed"))
}
