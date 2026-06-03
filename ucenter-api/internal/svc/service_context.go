// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"grpc-common/ucenter/ucclient"
	"ucenter-api/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	URegisterRpc ucclient.Register
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		URegisterRpc: ucclient.NewRegister(zrpc.MustNewClient(c.UcenterRpc)),
	}
}
