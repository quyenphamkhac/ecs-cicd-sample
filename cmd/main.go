package main

import (
	"github.com/quyenphamkhac/ecs-cicd-sample/server"
)

func main() {
	s := server.NewHttpServer()
	s.Run(":8080")
}
