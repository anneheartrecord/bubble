package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangstudy/gowebjinjie/bubble/dao"
	"golangstudy/gowebjinjie/bubble/model"
)

func IndexHandler(c *gin.Context)  {
		c.HTML(200,"index.html",nil)
}
func CreateTodo(c *gin.Context)  {
		var todo model.Todo
		c.BindJSON(&todo)
		if err:=dao.DB.Create(&todo).Error;err!=nil {
			fmt.Println("failed to insert db",err)
			c.JSON(200,gin.H{
				"code" :200,
				"msg": "failed to insert db",
			})
			return
		}
		c.JSON(200,gin.H{
			"code":200,
			"msg": "ok",
		})
}
func ShowTodo(c *gin.Context)  {
	var todoList []model.Todo
	if err:=dao.DB.Find(&todoList).Error;err!=nil {
		c.JSON(200,gin.H{
			"code":200,
			"msg": err.Error(),
		})
	}else {
		c.JSON(200,todoList)
	}
}
func UpdateTodo(c *gin.Context)  {
	id,ok:=c.Params.Get("id")
	if !ok {
		fmt.Println("wrong id")
		return
	}
	var todo model.Todo
	if err:=dao.DB.Where("id=?",id).First(&todo).Error;err!=nil {
		c.JSON(200,gin.H{
			"code":200,
			"msg": err.Error(),
		})
		return
	}
	c.ShouldBind(&todo)
	if err:=dao.DB.Save(&todo).Error;err!=nil {
		c.JSON(200,gin.H{
			"code": 200,
			"msg": err.Error(),
		})
	} else {
		c.JSON(200,todo)
	}
}
func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(200, gin.H{"msg": "无效id"})
		return
	}
	if err := dao.DB.Where("id=?", id).Delete(model.Todo{}).Error; err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"msg": "ok"})
	}
}