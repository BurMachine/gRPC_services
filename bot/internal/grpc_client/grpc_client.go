package grpcClient

import (
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func RunClient(port string, logger *zerolog.Logger) {
	conn, err := grpc.Dial(port)
	if err != nil {
		logger.Fatal().Err(err).Msg("grpc dial error")
	}
	println(conn)

}
