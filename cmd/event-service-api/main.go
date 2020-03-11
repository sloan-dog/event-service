package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"sloan.com/service/cmd/event-service-api/handlers"
)

func main() {
	if err := run(); err != nil {
		log.Println("error: ", err)
	}
}

func run() error {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	api := http.Server{
		Addr:    "0.0.0.0:3005",
		Handler: handlers.API(),
		// read timeout?
		// write timeout?
	}

	go func() {
		api.ListenAndServe() // todo handle errors
	}()

	select {
	case <-shutdown: // todo care about the signal
		// shut it down
		fmt.Printf("\nshutting down\n")
		ctx, cancel := context.WithTimeout(context.Background(), 5000)
		defer cancel()
		err := api.Shutdown(ctx) // give it a context?
		if err != nil {
			// whaterver
			err = api.Close()
		}
	}
	return nil
}
