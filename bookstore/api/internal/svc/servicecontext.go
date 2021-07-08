package svc

import (
	"bookstore/api/internal/config"
	"bookstore/services/user/userservice"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	User   userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userservice.NewUserService(zrpc.MustNewClient(c.User)),
	}
}
