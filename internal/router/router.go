package router

import (
	"github.com/gin-gonic/gin"
	"github.com/keshu12345/overlap/internal/handler"
	"github.com/keshu12345/overlap/service"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	svc := service.NewOverlapService()
	hdl := handler.NewOverlapHandler(svc)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/overlap-check", hdl.CheckOverlap)
	}

	return r
}
