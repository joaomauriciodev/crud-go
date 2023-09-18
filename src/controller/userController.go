package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomauriciodev/crud-go/src/configuration/validation"
	"github.com/joaomauriciodev/crud-go/src/controller/model/request"
)

func FindUserById(c *gin.Context) {}

func FindUserByEmail(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
	c.JSON(http.StatusOK, userRequest)

}

func UpdateUser(c *gin.Context) {}

func DeleteUser(c *gin.Context) {}
