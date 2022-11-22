package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"service-app/web"
)

// go get github.com/gorilla/mux
// go mod tidy - to remove unused dependency // to add dep

func Api() http.Handler {
	//r := mux.NewRouter()

	app := web.App{
		Router: mux.NewRouter(),
	}

	//r.HandleFunc("/check", check)
	app.HandleFunc(http.MethodGet, "/check", check)

	return app
	//return r
}

//func abc(writer http.ResponseWriter, request *http.Request) error {
//	ctx := request.Context()
//	user.Getuser(ctx)
//}
