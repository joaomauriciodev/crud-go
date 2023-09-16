package controller

import (
	"github.com/gin-gonic/gin"
	resterr "github.com/joaomauriciodev/crud-go/src/configuration/rest_err"
)

func FindUserById(c *gin.Context) {}

func FindUserByEmail(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	err := resterr.NewBadRequestError("Você chamou a rota errada")
	c.JSON(err.Code, err)
}

func UpdateUser(c *gin.Context) {}

func DeleteUser(c *gin.Context) {}
