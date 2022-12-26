package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/user/api/internal/config"
	userClient "mall/service/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	UserRpc userClient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userClient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
