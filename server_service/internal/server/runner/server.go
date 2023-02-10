package runner

import (
	"burmachine/internal/server/implemenatation"
	pb "burmachine/pkg/api"
	"google.golang.org/grpc"
	"net"
)

func Run(port string) error {
	grpcServ := grpc.NewServer()
	implementedServ := implemenatation.GrpcServer{}
	pb.RegisterFirstServicePbServer(grpcServ, &implementedServ)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	err = grpcServ.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
