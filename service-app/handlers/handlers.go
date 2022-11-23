package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"service-app/auth"
	"service-app/middlewares"
	"service-app/web"
)

// go get github.com/gorilla/mux
// go mod tidy - to remove unused dependency // to add dep

func Api(log *log.Logger, a *auth.Auth) http.Handler {
	//r := mux.NewRouter()

	app := web.App{
		Router: mux.NewRouter(),
	}
	m := middlewares.Mid{Log: log, A: a}

	//r.HandleFunc("/check", check)
	app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(m.Panic(m.Authenticate(m.Authorize(check, auth.RoleAdmin))))))

	return app
	//return r
}

//func abc(writer http.ResponseWriter, request *http.Request) error {
//	ctx := request.Context()
//	user.Getuser(ctx)
//}
