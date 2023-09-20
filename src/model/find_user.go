package model

import resterr "github.com/joaomauriciodev/crud-go/src/configuration/rest_err"

func (ud *UserDomain) FindUser(string) (*UserDomain, *resterr.RestErr) {
	return ud,nil
}