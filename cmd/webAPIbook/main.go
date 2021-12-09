package main

import (
	"fmt"
	"golang_ninja/webAPIbook/config"
	"golang_ninja/webAPIbook/pkg/transport"
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

	time.Sleep(time.Second * 5)

	err = transport.Server()
	if err != nil {
		log.Fatal("error program. stop app")
	}
}
