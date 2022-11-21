package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"mvc-app/model"
	"mvc-app/web"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	userIdString := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseUint(userIdString, 10, 64)

	if err != nil {
		log.Println(err)
		appErr := web.ApplicationError{Message: "please provide a valid user id"}
		//jsonData, _ := json.Marshal(appErr)
		//w.Write(jsonData)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(appErr) // it will write the json directly to the ResponseWriter
		return                            // this is important // don't forget to return
	}

	user, err := model.FetchUser(userId)
	if err != nil {
		log.Println(err)
		appErr := web.ApplicationError{Message: "user not found"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(appErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	return

}

func Home(w http.ResponseWriter, r *http.Request) {
	panic("in handler")
	fmt.Fprintln(w, "Hello this is home page")
}
