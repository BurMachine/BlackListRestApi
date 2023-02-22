package handlers

import (
	"log"
	"net/http"
)

func (h *Handlers) Addition(w http.ResponseWriter, r *http.Request) {
	log.Println("Addition")
}
