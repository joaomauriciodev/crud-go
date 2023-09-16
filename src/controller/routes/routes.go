package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaomauriciodev/crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:id", controller.FindUserById)
	//r.GET("/user/:email", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user", controller.UpdateUser)
	r.DELETE("/user", controller.DeleteUser)
}
