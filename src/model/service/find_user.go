package service

import (
	resterr "github.com/joaomauriciodev/crud-go/src/configuration/rest_err"
	"github.com/joaomauriciodev/crud-go/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface,
	*resterr.RestErr) {
	return nil,nil
}