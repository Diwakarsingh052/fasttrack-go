package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-app/handlers"
	"time"
)

func main() {

	err := startApp()
	if err != nil {
		log.Fatal(err)
	}

}

func startApp() error {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt) // this will notify the shutdown chan if someone presses ctr+c

	// providing config for the server
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		Handler:      handlers.Api(),
	}
	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("main: API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error %w", err)

	case sig := <-shutdown:
		log.Printf("main: %v : Start shutdown", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Asking listener to shut down and shed load.
		err := api.Shutdown(ctx) // first trying to cleanly shutdown
		if err != nil {
			err = api.Close() // forcing shutdown
			return fmt.Errorf("could not stop server gracefully %w", err)
		}
	}

	return nil
}
