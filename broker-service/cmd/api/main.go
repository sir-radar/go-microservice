package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "3000"

type Config struct {}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", PORT)

	// define http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
