package handlers

import (
	"blacklistApi/internal/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func (h *Handlers) Search(w http.ResponseWriter, r *http.Request) {
	number := r.URL.Query().Get("number")
	name := r.URL.Query().Get("name")
	if number == "" && name == "" {
		log.Println("no params to search")
		http.Error(w, errors.New("empty number").Error(), http.StatusBadRequest)
		return
	}
	var users []models.Addiction
	var err error
	if number != "" {
		users, err = h.storage.Search("phone_number", number)
		if err != nil {
			log.Println("search error")
			http.Error(w, errors.New("search error").Error(), http.StatusBadRequest)
			return
		}
	} else {
		users, err = h.storage.Search("user_name", name)
		if err != nil {
			log.Println("search error")
			http.Error(w, errors.New("search error").Error(), http.StatusBadRequest)
			return
		}
	}
	resp, err := json.Marshal(users)
	if err != nil {
		log.Println("resp marshalling error")
		http.Error(w, errors.New("search error").Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
