package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting Broker Service on port %s", webPort)

	// setup http server
	svc := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start server
	err := svc.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
