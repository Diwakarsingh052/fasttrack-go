package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"service-app/auth"
	"service-app/web"
	"time"
)

type Mid struct {
	Log *log.Logger
	A   *auth.Auth
}

func (m *Mid) Logger(next web.HandlerFunc) web.HandlerFunc {

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		v, ok := ctx.Value(web.KeyValue).(*web.Values) // type assertion and making sure values struct with the trace id is present
		if !ok {
			return fmt.Errorf("web.Values missing from the context")
		}
		m.Log.Printf("%s : started   : %s %s ",
			v.TraceId,
			r.Method, r.URL.Path)

		err := next(ctx, w, r) // executing the next handlerFunc or the middleware in the chain

		m.Log.Printf("%s : completed : %s %s ->  (%d)  (%s)",
			v.TraceId,
			r.Method, r.URL.Path,
			v.StatusCode, time.Since(v.Now),
		)

		return err

	}
}
