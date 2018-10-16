package main

import (
	"fmt"

	"github.com/mic0331/go-rabitmq-distributed-app/distributed/coordinator"
)

func main() {
	ql := coordinator.NewQueueListener()
	go ql.ListenForNewSource()

	var a string
	fmt.Scanln(&a)
}
