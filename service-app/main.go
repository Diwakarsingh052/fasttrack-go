package main

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-app/auth"
	"service-app/data/users"
	"service-app/database"
	"service-app/handlers"
	"time"
)

func main() {

	l := log.New(os.Stdout, "users :", log.LstdFlags)

	err := startApp(l)
	if err != nil {
		log.Fatal(err)
	}

}

func startApp(log *log.Logger) error {

	// =========================================================================
	// Start Database
	log.Println("main : Started : Initializing db support")
	db, err := database.Open()
	if err != nil {
		return fmt.Errorf("connecting to db %w", err)
	}
	if err := db.Ping(); err != nil {
		return err
	}
	uDB := &users.DbService{DB: db}

	// =========================================================================
	// Initialize authentication support
	log.Println("main : Started : Initializing authentication support")
	privatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		return fmt.Errorf("reading auth private key %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("parsing auth private key %w", err)
	}
	a, err := auth.NewAuth(privateKey)
	if err != nil {
		return fmt.Errorf("constructing auth %w", err)
	}

	//Initialize server
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt) // this will notify the shutdown chan if someone presses ctr+c

	// providing config for the server
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		Handler:      handlers.Api(log, a, uDB),
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
