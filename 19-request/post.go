package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//http.Post()
	user := map[string]string{"first_name": "raj"}
	u, _ := json.Marshal(user)

	// this is the construction of the req // we haven't made the request yet
	req, err := http.NewRequest(http.MethodPost, "http://httpbin.org/post", bytes.NewReader(u))

	req.Header.Set("Content-Type", "application/json") // setting headers // we will send a json to the server

	resp, err := http.DefaultClient.Do(req) // doing the request

	// this error means if we can't call the remote server
	if err != nil {
		log.Fatalln(err)
	}

	data, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data))
}
