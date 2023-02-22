package handlers

import (
	"log"
	"net/http"
)

func (h *Handlers) Removal(w http.ResponseWriter, r *http.Request) {
	log.Println("Addition")
}
