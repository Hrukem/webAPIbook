package main

import (
	"fmt"
	"golang_ninja/webAPIbook/pkg/config"
	"golang_ninja/webAPIbook/pkg/server"
	"log"
)

func main() {
	fmt.Println("start application")

	err := config.Config()
	if err != nil {
		log.Fatal(err)
	}

	server.Server()
}
