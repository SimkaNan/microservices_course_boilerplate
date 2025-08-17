package api

import (
	"context"
	pb "github.com/olezhek28/microservices_course_boilerplate/gen/chat-server/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type server struct {
	pb.UnimplementedChatServerServer
}

// NewServer создаёт и возвращает gRPC-обработчик сервиса <название>
// с зарегистрированными зависимостями/интерсепторами.
// Используется из cmd/main.go для регистрации в grpc.Server.
func NewServer() *server {
	return &server{
		pb.UnimplementedChatServerServer{},
	}
}

func (s *server) Create(_ context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	users := in.GetUsernames()
	if users == nil {
		err := status.Error(codes.InvalidArgument, "empty fields")
		log.Println(err)
		return nil, err
	}
	for _, user := range users {
		log.Println(user)
	}
	return &pb.CreateResponse{}, nil
}

func (s *server) Delete(_ context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	id := in.GetId()
	log.Println(id)
	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(_ context.Context, in *pb.SendMessageRequest) (*emptypb.Empty, error) {
	log.Println(in)
	return &emptypb.Empty{}, nil
}
