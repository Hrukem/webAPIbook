package main

import (
	"fmt"
	"golang_ninja/webAPIbook/pkg/config"
	"golang_ninja/webAPIbook/pkg/transport"
	"log"
	"time"
)

func main() {
	fmt.Println("start application")
	time.Sleep(time.Second * 5)

	err := config.Config()
	if err != nil {
		log.Fatal(err)
	}

	transport.Server()
}
