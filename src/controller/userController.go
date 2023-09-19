package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomauriciodev/crud-go/src/configuration/logger"
	"github.com/joaomauriciodev/crud-go/src/configuration/validation"
	"github.com/joaomauriciodev/crud-go/src/controller/model/request"
	"go.uber.org/zap"
)

func FindUserById(c *gin.Context) {}

func FindUserByEmail(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
		zap.String("methodHttp", "POST"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		logger.Error("Error trying to validate user info", err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
	c.JSON(http.StatusOK, userRequest)

}

func UpdateUser(c *gin.Context) {}

func DeleteUser(c *gin.Context) {}
