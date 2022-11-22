package controllers

import (
	"encoding/json"
	"log"
	"mvc-app/model"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// this line set your  ContentType as json
	w.Header().Set("Content-Type", "application/json")
	userIdString := r.URL.Query().Get("user_id")

	usedId, err := strconv.ParseUint(userIdString, 10, 64)

	// this error happens when you have user id in wrong format and not we can't convert it to uint
	if err != nil {
		jsonErr := map[string]string{"Message": err.Error()}

		w.WriteHeader(http.StatusBadRequest) // setting error status code
		json.NewEncoder(w).Encode(jsonErr)   // writing json back to the client
		log.Println(err)
		return

	}

	user, err := model.FetchUser(usedId)
	//case : user not found
	if err != nil {
		jsonErr := map[string]string{"Message": err.Error()}
		w.WriteHeader(http.StatusBadRequest) // setting error status code
		json.NewEncoder(w).Encode(jsonErr)
		log.Println(err)
		return
	}

	// case : user found
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
	return

}
