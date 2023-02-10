package main

import (
	"burmachine/internal/server/runner"
	"log"
)

func main() {
	err := runner.Run(":8081")
	if err != nil {
		log.Fatal("failed grpc")
	}
}
