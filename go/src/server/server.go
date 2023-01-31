package server

import (
	"github.com/gin-gonic/gin"
	vocabulary "nothing-behind.com/sample_gin/controller"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	v := r.Group("/vocabularies")
	{
		ctrl := vocabulary.Controller{}
		v.GET("/", ctrl.Index)
	}

	return r
}
