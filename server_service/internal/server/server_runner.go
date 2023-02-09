package server

import "context"

type GrpcServer struct {
}

func (s *GrpcServer) Add(context.Context, *RequestMessage) (*ResponseMessage, error) {

}
