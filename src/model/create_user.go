package model

import (
	"fmt"

	"github.com/joaomauriciodev/crud-go/src/configuration/logger"
	resterr "github.com/joaomauriciodev/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *resterr.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}