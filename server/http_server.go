package server

import (
	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/ecs-cicd-sample/controller"
	"github.com/quyenphamkhac/ecs-cicd-sample/pkg/middlewares"
)

type httpServer struct{}

func NewHttpServer() *httpServer {
	return &httpServer{}
}

func (s *httpServer) Run(addr string) {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.ErrorsMiddleware(gin.ErrorTypeAny))

	heatlhCtrl := controller.NewHealthController()

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		{
			health.GET("", heatlhCtrl.Health)
		}
	}
	r.Run(addr)
}
