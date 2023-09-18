package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	resterr "github.com/joaomauriciodev/crud-go/src/configuration/rest_err"
	"github.com/joaomauriciodev/crud-go/src/controller/model/request"
)

func FindUserById(c *gin.Context) {}

func FindUserByEmail(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := resterr.NewBadRequestError(fmt.Sprintf("There are some incorret fields, error=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
	c.JSON(http.StatusOK, userRequest)

}

func UpdateUser(c *gin.Context) {}

func DeleteUser(c *gin.Context) {}
