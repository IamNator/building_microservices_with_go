package main

import (
	"github.com/IamNator/building_microservices_with_go/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main(){

	 l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGood(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/good", gh)

	s := &http.Server{
		Addr: ":8009",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 2*time.Second,
		WriteTimeout: 2*time.Second,
	}

	log.Println("Server running @localhost:8009")
	_ = s.ListenAndServe()

}
