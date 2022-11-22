package web

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Values struct {
	TraceId    string
	Now        time.Time
	StatusCode int
}

// creating custom type for storing the value in the context
type ctxKey int

const KeyValue ctxKey = 1

type App struct {
	*mux.Router
}

// http.HandleFunc("/",handleFunc)
// web.HandleFunc("GET","/",handleFunc)

// HandlerFunc is a custom type like http.HandlerFunc func in standard library
type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func (a *App) HandleFunc(method string, path string, handler HandlerFunc) {

	// wrapping the actual call in a variable which have a same signature as a handlerFunc ,
	//so we can pass it to gorilla mux to process it
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		v := Values{
			TraceId: uuid.NewString(),
			Now:     time.Now(),
		}
		ctx = context.WithValue(ctx, KeyValue, v)

		err := handler(ctx, w, r)
		if err != nil {
			log.Println(err)
			return
		}
	}
	// mux router can accept the f var because the signature of h var matches to func(w http.ResponseWriter, r *http.Request)
	a.Router.HandleFunc(path, f).Methods(method) // executing the handler func

}
