package main

import (
	"log/slog"
	"net"

	"github.com/aakash19here/grpc_exam/proto/generated/exampb"
	"github.com/aakash19here/grpc_exam/server/servers"
	"github.com/aakash19here/grpc_exam/utils"
	"google.golang.org/grpc"
)

func main() {
	utils.InitLogger(true)

	// listens to tcp connections at port 50051
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		slog.Error("failed to listen", "error", err)
	}

	s := grpc.NewServer()

	//register services
	// since servers.NewExamServerService() implements UnimplementedExamServiceServer we can pass it to register exam service
	exampb.RegisterExamServiceServer(s, servers.NewExamServerService())

	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve", "error", err)
	}
}
