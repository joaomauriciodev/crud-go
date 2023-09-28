package service

import (
	resterr "github.com/joaomauriciodev/crud-go/src/configuration/rest_err"
	"github.com/joaomauriciodev/crud-go/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {

}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *resterr.RestErr
	UpdateUser(string,model.UserDomainInterface) *resterr.RestErr
	FindUser(string,) (*model.UserDomainInterface, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}
