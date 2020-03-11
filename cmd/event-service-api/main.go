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
	"sloan.com/service/internal/constants"
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
		Addr:    fmt.Sprintf("%s:%s", "0.0.0.0", constants.Port),
		Handler: handlers.API(),
		// read timeout?
		// write timeout?
	}

	go func() {
		api.ListenAndServe()
	}()

	select {
	case <-shutdown:
		// shut it down
		fmt.Printf("\n<===== shutting down =====>\n")
		ctx, cancel := context.WithTimeout(context.Background(), 5000)
		defer cancel()
		err := api.Shutdown(ctx)
		if err != nil {
			err = api.Close()
		}
	}
	return nil
}
