package implemenatation

import (
	pb "burmachine/pkg/api"
	"context"
)

type GrpcServer struct {
	pb.UnimplementedFirstServicePbServer
}

func (s *GrpcServer) mustEmbedUnimplementedFirstServicePbServer() {
	//TODO implement me
	panic("implement me")
}

func (s *GrpcServer) Add(ctx context.Context, req *pb.RequestMessage) (*pb.ResponseMessage, error) {
	return &pb.ResponseMessage{Result: req.GetX() + req.GetY()}, nil
}
