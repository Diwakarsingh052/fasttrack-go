package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	//http.Post()
	user := map[string]string{"first_name": "raj"}
	u, _ := json.Marshal(user)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
	defer cancel()

	// this is the construction of the req // we haven't made the request yet
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://httpbin.org/post", bytes.NewReader(u))

	req.Header.Set("Content-Type", "application/json") // setting headers // we will send a json to the server

	resp, err := http.DefaultClient.Do(req) // doing the request // if timeout happens we will get an error

	// this error means if we can't call the remote server
	if err != nil {
		log.Fatalln(err)
	}

	data, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data))
}
