package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"service-app/auth"
	"service-app/data/users"
	"service-app/middlewares"
	"service-app/web"
)

// go get github.com/gorilla/mux
// go mod tidy - to remove unused dependency // to add dep

func Api(log *log.Logger, a *auth.Auth, uDB *users.DbService) http.Handler {
	//r := mux.NewRouter()

	app := web.App{
		Router: mux.NewRouter(),
	}
	uh := userHandlers{
		DbService: uDB,
		Auth:      a,
	}
	m := middlewares.Mid{Log: log, A: a}

	//r.HandleFunc("/check", check)
	app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(m.Panic(m.Authenticate(m.Authorize(check, auth.RoleAdmin))))))
	app.HandleFunc(http.MethodPost, "/signup", m.Logger(m.Error(m.Panic(uh.SignUp))))
	app.HandleFunc(http.MethodPost, "/login", m.Logger(m.Error(m.Panic(uh.Login))))
	app.HandleFunc(http.MethodGet, "/add", m.Logger(m.Error(m.Panic(m.AuthenticateCookie(uh.AddInventory)))))
	app.HandleFunc(http.MethodGet, "/view", m.Logger(m.Error(m.Panic(m.AuthenticateCookie(uh.ViewInventory)))))
	return app
	//return r
}

//func abc(writer http.ResponseWriter, request *http.Request) error {
//	ctx := request.Context()
//	user.Getuser(ctx)
//}
