package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"service-app/auth"
	"service-app/data/users"
	"service-app/web"
)

// AddInventory allow user to add some products in db
func (h *userHandlers) AddInventory(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, ok := ctx.Value(web.KeyValue).(*web.Values)
	if !ok {
		return fmt.Errorf("web.Values missing from the context")
	}

	claims, ok := ctx.Value(auth.Key).(auth.Claims)

	if !ok {
		return web.NewRequestError(errors.New("not authenticated"), http.StatusUnauthorized)
	}

	var newInv users.NewShirtInventory
	err := web.Decode(r, &newInv)
	if err != nil {
		return err
	}
	s, err := h.CreateInventory(ctx, newInv, claims.Subject, v.Now)
	if err != nil {
		return web.NewRequestError(errors.New("problem in creating inventory"), http.StatusBadRequest)
	}

	web.Respond(ctx, w, s, http.StatusCreated)
	return nil

}

// ViewInventory helps to view all the inventory for the current user
func (h *userHandlers) ViewInventory(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims, ok := ctx.Value(auth.Key).(auth.Claims)

	if !ok {
		return web.NewRequestError(errors.New("not authenticated"), http.StatusUnauthorized)
	}

	shirts, err := h.ViewAll(ctx, claims.Subject)

	if err != nil {
		return web.NewRequestError(errors.New("problem in viewing inventory"), http.StatusBadRequest)
	}

	web.Respond(ctx, w, shirts, http.StatusFound)
	return nil

}
