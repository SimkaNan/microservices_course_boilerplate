package api

import (
	"context"
	"fmt"
	pb "github.com/olezhek28/microservices_course_boilerplate/chat-server/pkg/noteV1"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type server struct {
	pb.UnimplementedChat_ServerServer
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	users := in.GetUsernames()
	if users == nil {
		log.Println(fmt.Errorf("empty fields"))
		return nil, fmt.Errorf("empty fields")
	}
	for _, user := range users {
		log.Println(user)
	}
	return nil, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	id := in.GetId()
	log.Println(id)
	return nil, nil
}

func (s *server) SendMessage(ctx context.Context, in *pb.SendMessageRequest) (*emptypb.Empty, error) {
	log.Println(in)
	return nil, nil
}
