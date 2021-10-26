package main

import (
	//"context"
	//"fmt"
	"log"
	"net"

	//t "time"

	chittychat "github.com/Stalin1999/chatfunction/chittychat"

	"google.golang.org/grpc"
)

type Server struct {
	chittychat.UnimplementedServiceServer
}

/* func (s *Server) GetTime(ctx context.Context, in *chittychat.GetTimeRequest) (*chittychat.GetTimeReply, error) {
	fmt.Printf("Received get time request")
	return &chittychat.GetTimeReply{Reply: t.Now().String()}, nil
} */

func main() {
	// init listener and tcp on port 9080
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Could not listen to port 9080: %v", err)
	}

	//gRPC server instance
	grpcServer := grpc.NewServer()

	//cc := chittychat.ChatServer{}
	chittychat.RegisterServiceServer(grpcServer, &Server{})

	//gRPC listen and serve
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}
