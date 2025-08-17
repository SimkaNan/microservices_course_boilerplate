package api

import (
	"context"
	pb "github.com/olezhek28/microservices_course_boilerplate/gen/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type server struct {
	pb.UnimplementedUserAPIServer
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	u := in.GetUser()
	if u == nil {
		err := status.Error(codes.InvalidArgument, "empty fields")
		log.Println(err)
		return nil, err
	}
	log.Println(u)
	return &pb.CreateResponse{}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	id := in.GetId()
	if id == 0 {
		err := status.Error(codes.NotFound, "user not found")
		log.Println(err)
		return nil, err
	}
	log.Println(id)
	return &pb.GetResponse{}, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateRequest) (*emptypb.Empty, error) {
	u := in.GetUser()
	if u == nil {
		err := status.Error(codes.InvalidArgument, "empty fields")
		log.Println(err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	id := in.GetId()
	if id == 0 {
		err := status.Error(codes.NotFound, "user not found")
		log.Println(err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
