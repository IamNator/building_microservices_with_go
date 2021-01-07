package handlers

import (
	"log"
	"net/http"
	"encoding/json"
)

type Good struct {
	l *log.Logger
}

func NewGood(l *log.Logger) *Good {
	return &Good{l}
}

func (g *Good) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	resp := struct {
		Message string `json:"message"`
	}{
		"Great",
	}

	rw.Header().Set("content-type", "application/json")
	json.NewEncoder(rw).Encode(resp)
}

