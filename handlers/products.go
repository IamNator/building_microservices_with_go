package handlers

import (
	jsonWriter "github.com/IamNator/JsonWrite"
	"github.com/IamNator/building_microservices_with_go/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			jsonWriter.Error(rw, "Invalid URL", http.StatusRequestURITooLong)
			return
		}

		if len(g[0]) != 1 {
			jsonWriter.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			jsonWriter.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		p.l.Println(id, " received")

	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProduct(rw http.ResponseWriter, _ *http.Request){
	lp := data.GetProduct()
	err := lp.ToJson(rw)
	if err!=nil {
		jsonWriter.Error(rw, "Unable to Marshall Json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r * http.Request){
	prod := &data.Product{}

	err := prod.FromJson(r.Body)
	if err != nil {
		jsonWriter.Error(rw, "Unable to Unmarshal Json", http.StatusBadRequest)
		return
	}
	data.AddProduct(prod)
}