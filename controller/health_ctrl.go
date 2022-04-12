package controller

import "github.com/gin-gonic/gin"

type healthCtrl struct {
}

func NewHealthController() *healthCtrl {
	return &healthCtrl{}
}

func (c *healthCtrl) Health(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}
