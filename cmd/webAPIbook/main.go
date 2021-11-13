package main

import (
	"fmt"
	"golang_ninja/webAPIbook/config"
	"golang_ninja/webAPIbook/pkg/transport"
	"log"
	"time"
)

func main() {
	fmt.Println("start application")

	err := config.Config()
	if err != nil {
		log.Fatal("error reade config", err)
	}

	time.Sleep(time.Second * 5)

	err = transport.Server()
	if err != nil {
		log.Fatal("error program. stop app")
	}
}
