package handlers

import (
	"context"
	"net/http"
	"service-app/web"
)

func check(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{Status: "ok"}
	_ = status
	//panic("i need to panic")
	return web.Respond(ctx, w, status, http.StatusOK)
	//return errors.New("something went wrong in the db")
	//return web.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)

}

//v1
//func check(w http.ResponseWriter, r *http.Request) {
//	status := struct {
//		Status string
//	}{Status: "ok"}
//
//	json.NewEncoder(w).Encode(status)
//}
