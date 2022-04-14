package main

import (
	"fmt"

	"github.com/quyenphamkhac/ecs-cicd-sample/server"
)

func main() {
	fmt.Println("Hello World")
	s := server.NewHttpServer()
	s.Run(":8080")
}
