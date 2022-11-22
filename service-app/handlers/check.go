package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

func check(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{Status: "ok"}

	return json.NewEncoder(w).Encode(status)

}

//v1
//func check(w http.ResponseWriter, r *http.Request) {
//	status := struct {
//		Status string
//	}{Status: "ok"}
//
//	json.NewEncoder(w).Encode(status)
//}
