package api

import (
	"context"
	"fmt"
	pb "github.com/olezhek28/microservices_course_boilerplate/auth/pkg/noteV1"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type server struct {
	pb.UnimplementedUserAPIServer
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	u := in.GetUser()
	if u == nil {
		log.Println(fmt.Errorf("empty fields"))
		return nil, fmt.Errorf("empty fields")
	}
	log.Println(u)
	return nil, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	id := in.GetId()
	if id == 0 {
		log.Println(fmt.Errorf("user not found"))
		return nil, fmt.Errorf("user not found")
	}
	log.Println(id)
	return nil, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateRequest) (*emptypb.Empty, error) {
	u := in.GetUser()
	if u == nil {
		log.Println(fmt.Errorf("empty fields"))
		return nil, fmt.Errorf("empty fields")
	}
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	id := in.GetId()
	if id == 0 {
		log.Println(fmt.Errorf("user not found"))
		return nil, fmt.Errorf("user not found")
	}
	return &emptypb.Empty{}, nil
}
