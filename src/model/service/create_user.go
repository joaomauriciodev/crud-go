package service

import (
	"fmt"

	"github.com/joaomauriciodev/crud-go/src/configuration/logger"
	resterr "github.com/joaomauriciodev/crud-go/src/configuration/rest_err"
	"github.com/joaomauriciodev/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *resterr.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(ud)

	return nil
}