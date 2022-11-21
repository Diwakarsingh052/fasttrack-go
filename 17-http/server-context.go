package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/home", randomData)
	http.ListenAndServe(":8080", nil)
}

func randomData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("started")
	defer log.Println("ended")

	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "random"+
			" data")
	case <-ctx.Done(): // it will tell when a ctx is cancelled or finished
		err := ctx.Err()
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// try to press ctr+c to end the request in between to see cancellation in action
}
