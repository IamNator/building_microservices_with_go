package main

import (
	"context"
	"github.com/IamNator/building_microservices_with_go/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){

	 l := log.New(os.Stdout, "product-api", log.LstdFlags)


	ph := handlers.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)


	s := &http.Server{
		Addr: ":8009",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 2*time.Second,
		WriteTimeout: 2*time.Second,
	}

	go func(){
		l.Println("Server running @localhost:8009")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err.Error())
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	t := time.Now().Add(30*time.Second)
	tc,_ := context.WithDeadline(context.Background(), t)
	s.Shutdown(tc)

}
