package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"restar/configs"
	"restar/pkg/user/pb"
	"time"
)

func main() {
	conn, err := grpc.Dial(configs.NewConfig().Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msgf("cant connect to grpc service %s", configs.NewConfig().Host)
	}

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.UserInfo(ctx, &pb.UserRequest{Id: 23})
	if err != nil {
		log.Err(err).Send()
	}

	log.Printf("user response: %v", resp.Name)
}
