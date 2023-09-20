package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomauriciodev/crud-go/src/configuration/logger"
	"github.com/joaomauriciodev/crud-go/src/configuration/validation"
	"github.com/joaomauriciodev/crud-go/src/controller/model/request"
	"github.com/joaomauriciodev/crud-go/src/model"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func FindUserById(c *gin.Context) {}

func FindUserByEmail(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		logger.Error("Error trying to validate user info", err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email, 
		userRequest.Password, 
		userRequest.Name, 
		userRequest.Age)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey","createUser"))

	c.String(http.StatusOK, "")

}

func UpdateUser(c *gin.Context) {}

func DeleteUser(c *gin.Context) {}
