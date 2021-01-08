package handlers

import (
	jsonWriter "github.com/IamNator/JsonWrite"
	"github.com/IamNator/building_microservices_with_go/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		p.getProduct(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProduct(rw http.ResponseWriter, r *http.Request){
	lp := data.GetProduct()
	err := lp.ToJson(rw)
	if err!=nil {
		jsonWriter.Error(rw, "Unable to Marshall Json", http.StatusInternalServerError)
	}
}