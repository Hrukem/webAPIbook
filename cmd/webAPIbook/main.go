package main

import (
	"fmt"
	"golang_ninja/webAPIbook/pkg/config"
	"golang_ninja/webAPIbook/pkg/storage/mongoDB"
	"golang_ninja/webAPIbook/pkg/storage/postgress"
	"golang_ninja/webAPIbook/pkg/transport/g_r_p_c/grpcClient"
	"golang_ninja/webAPIbook/pkg/transport/g_r_p_c/grpcServer"
	"golang_ninja/webAPIbook/pkg/transport/h_t_t_p"
	"log"
	"time"
)

// @title Swagger WebApiBook
// version 1.0
// description Service for storing data about books

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4004
// @BasePath /

func main() {
	fmt.Println("start application")

	err := config.Config()
	if err != nil {
		log.Fatal("error read config: ", err)
	}

	dbPostgres, err := postgress.NewDb()
	if err != nil {
		log.Fatal("error create PostgresDb")
	}

	mngdbCollection, err := mongoDB.NewMongoDB()
	if err != nil {
		log.Fatal("error create MongoDB")
	}

	time.Sleep(time.Second * 5)

	go grpcServer.RunGRPCserver(mngdbCollection)

	loggingInMongo := make(chan grpcServer.L, 10)
	go grpcClient.RunGRPCclient(loggingInMongo)

	err = h_t_t_p.Server(dbPostgres, loggingInMongo)
	if err != nil {
		log.Fatal("error program. stop app")
	}
}
