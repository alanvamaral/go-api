package model

import (
	"fmt"

	"github.com/alanvamaral/go-api/src/configuration/logger"
	"github.com/alanvamaral/go-api/src/configuration/rest_err"
	"github.com/alanvamaral/go-api/src/model"
	"go.uber.org/zap"
)

var (
	sUserDomainInterface model.UserDomainInterface
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser Model", zap.String("journey", "createUser"))

	ud.EncryptPassword()
	fmt.Println(ud)

	return nil
}
