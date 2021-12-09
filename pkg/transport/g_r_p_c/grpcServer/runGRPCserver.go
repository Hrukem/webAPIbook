package grpcServer

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type serv struct {
	mng *mongo.Collection
}

type L struct {
	Id  int64
	Log string
}

func (s *serv) mustEmbedUnimplementedReceiveLogsServer() {
	panic("implement me")
}

func (s *serv) Receive(ctx context.Context, req *LogRequest) (*Response, error) {
	logs := bson.D{{"id", req.GetId()}, {"logs", req.GetLog()}}

	_, err := s.mng.InsertOne(ctx, logs)
	if err != nil {
		log.Println("error insert in MongoDB in grpcServer.Receive()", err)
	}

	return &Response{Resp: "Ok, log receive"}, nil
}

func RunGRPCserver(mngCollection *mongo.Collection) {
	s := grpc.NewServer()
	srv := &serv{mngCollection}
	RegisterReceiveLogsServer(s, srv)

	l, err := net.Listen("tcp", ":4005")
	if err != nil {
		log.Println("error start network ", err)
	}
	log.Println("start grpc grpcServer on port 4005")

	err = s.Serve(l)
	if err != nil {
		log.Println("error start grpc grpcServer ", err)
	}
}
