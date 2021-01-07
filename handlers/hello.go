package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	resp := struct{
		Message string `json:"message"`
	}{
		"Hello",
	}

	rw.Header().Set("Content-Type", "application/json")
	err  := json.NewEncoder(rw).Encode(resp)
	if err != nil {
		h.l.Println(err.Error())
	}
}