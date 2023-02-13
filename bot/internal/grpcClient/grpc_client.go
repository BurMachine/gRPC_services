package grpcClient

import (
	pb "burmachineBot/pkg/api"
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func RunClient(port string, logger *zerolog.Logger, grpcChan chan int) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		logger.Fatal().Err(err).Msg("grpc dial error")
	}
	client := pb.NewFirstServicePbClient(conn)
	for {
		select {
		case <-grpcChan:
			a, err := client.Add(context.Background(), &pb.RequestMessage{X: 1, Y: 2})
			if err != nil {
				logger.Info().Err(err).Msg("error client")
			}
			fmt.Println(a)
		}
	}
}
