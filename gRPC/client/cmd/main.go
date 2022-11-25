package main

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-client/logs"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/do", LogViaGrpc)
	http.ListenAndServe(":8080", nil)
}

func LogViaGrpc(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}
	b, _ := io.ReadAll(r.Body)
	json.Unmarshal(b, &req)

	grpcService(req.Name, req.Data)

}

func grpcService(name, data string) {
	conn, err := grpc.Dial(":5000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)

	}
	defer conn.Close()
	c := logs.NewLogServiceClient(conn)

	l := &logs.LogRequest{LogEntry: &logs.Log{
		Name: name,
		Data: data,
	}}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := c.WriteLog(ctx, l)

	if err != nil {

		log.Println("we messed this up")
		return
	}

	log.Println("we did some grpc stuff", res)
}
