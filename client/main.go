package main

import (
	"log/slog"
	"os"

	"github.com/aakash19here/grpc_exam/client/clients"
	"github.com/aakash19here/grpc_exam/proto/generated/exampb"
	"github.com/aakash19here/grpc_exam/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	utils.InitLogger(true)

	if len(os.Args) < 2 {
		slog.Error("Usage: go run client/main.go [unary, server, client, bidi]")
		return
	}

	// create connection with the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error("Failed to conenct to the server", "error", err)
		return
	}

	// close the connection to release memory
	defer conn.Close()

	client := exampb.NewExamServiceClient(conn)

	switch os.Args[1] {
	case "unary":
		clients.Unary(client)
	case "server":
		clients.Server_stream(client)
	case "client":
		//call unary here
	case "bidi":
		//call unary here
	}
}
