package grpcClient

import (
	"context"
	"golang_ninja/webAPIbook/pkg/transport/g_r_p_c/grpcServer"
	"google.golang.org/grpc"
	"log"
)

func RunGRPCclient(loggingInMongo <-chan grpcServer.L) {
	conn, err := grpc.Dial(":4005", grpc.WithInsecure())
	if err != nil {
		log.Println("error create conn grpc in grpcClient.RunClient()", err)
	}

	client := grpcServer.NewReceiveLogsClient(conn)
	log.Println("start grpcClient on 4005")

	for {
		l := <-loggingInMongo

		answerGrpcServer, err :=
			client.Receive(
				context.Background(),
				&grpcServer.LogRequest{Id: l.Id, Log: l.Log},
			)
		if err != nil {
			log.Println("error receive log in MongoDB in grpcClient.RunGRPCclient()", err)
		}

		log.Println(answerGrpcServer)
	}
}
