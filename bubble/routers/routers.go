package routers

import (
	"github.com/gin-gonic/gin"
	"golangstudy/gowebjinjie/bubble/controller"
)

func SetupRouters() *gin.Engine {
	r:=gin.Default()
	r.Static("./static","static")
	r.LoadHTMLGlob("./templates/*")
	r.GET("/",controller.IndexHandler)
	//增删改查
	v1:=r.Group("v1")
	{
		v1.POST("/todo",controller.CreateTodo)
		v1.GET("/todo",controller.ShowTodo)
		v1.PUT("/todo/:id",controller.UpdateTodo)
		v1.DELETE("/todo/:id",controller.DeleteTodo)

	}
	return r
}

