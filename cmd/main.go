package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"garlic-mayo/pkg/common/router"
)

func main() {
	listenAddress := strings.Trim(os.Getenv("APP_LISTENER_ADDR"), " ")
	if listenAddress == "" {
		log.Fatal("environment variable APP_LISTENER_ADDR not set")
	}

	common.Start(listenAddress)
	fmt.Println("Server listening on: ", listenAddress)
	select {}
}
