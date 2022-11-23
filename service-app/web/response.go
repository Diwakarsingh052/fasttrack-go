package web

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func Respond(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {
	v, ok := ctx.Value(KeyValue).(*Values)
	if !ok {
		return fmt.Errorf("web.Values missing from the context")
	}

	v.StatusCode = statusCode

	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)

	return nil

}

func RespondError(ctx context.Context, w http.ResponseWriter, err error) error {

	var webErr *Error
	ok := errors.As(err, &webErr)
	// if webErr is being used to construct an Error then I will send the content of it to the end user
	if ok {
		er := ErrorResponse{Error: webErr.Err.Error()}
		err := Respond(ctx, w, er, webErr.Status)
		if err != nil {
			return err
		}
		return nil // it means RespondError got success in sending an error response back to the client
	}

	// if error is not of type webErr then I will not leak any internal details about app and give a generic error message
	er := ErrorResponse{Error: http.StatusText(http.StatusInternalServerError)}
	err = Respond(ctx, w, er, http.StatusInternalServerError)
	if err != nil {
		return err
	}
	return nil

}

func Decode(r *http.Request, val any) error {

	err := json.NewDecoder(r.Body).Decode(&val)

	if err != nil {
		return err
	}

	return nil
}
