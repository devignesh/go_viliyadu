package router

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/todo", controller.CreateTodopost)
	router.GET("/todos", controller.GetAllTodo)
	router.GET("/todo/:id", controller.GetTodo)
	router.PUT("/todo/:id", controller.UpdateTodo)
	router.DELETE("/todo/:id", controller.DeleteTodo)

}
