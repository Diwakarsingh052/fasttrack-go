package main

import (
	"context"
	"database/sql"
	"google.golang.org/grpc"
	"grpc-server/logs"
	"log"
	"net"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer //UnimplementedLogServiceServer must be embedded to have forward compatible implementations.
	*sql.DB                            // add dep for the gRPC server
}

const gRPCPort = "5000"

func main() {

	gRPCServer()
	//go gRPCServer() when gRPC sever would be run alongside http server
	//http.ListenAndServe(":8080", nil)
}

func gRPCServer() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalln(err)
	}
	//NewServer creates a gRPC server which has no service registered and has not started to accept requests yet
	s := grpc.NewServer()
	db, _ := sql.Open("", "")
	logs.RegisterLogServiceServer(s, &LogServer{DB: db})
	log.Printf("gRPC Server started on port %s", gRPCPort)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry() // input is the struct type log

	//do your business logic
	log.Println(input, "yes we used gRPC to log the message in our terminal")

	res := &logs.LogResponse{Result: "all good bro here"}

	return res, nil

}

func (l *LogServer) AccessTheDb() {

}
