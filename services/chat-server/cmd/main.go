package main

import (
	pb "github.com/olezhek28/microservices_course_boilerplate/gen/chat-server/v1"
	api "github.com/olezhek28/microservices_course_boilerplate/services/chat-server/internal/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterChatServerServer(s, api.NewServer())

	log.Println("server start")

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %w", err)
	}
}
