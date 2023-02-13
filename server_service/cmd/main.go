package main

import (
	"burmachine/internal/server/runner"
	"log"
)

func main() {
	err := runner.Run(":9090")
	if err != nil {
		log.Fatal("failed grpc")
	}
}
