// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"grpc-common/market/mclient"
	"market-api/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	ExchangeRateRpc mclient.ExchangeRate
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ExchangeRateRpc: mclient.NewRate(zrpc.MustNewClient(c.MarketRpc)),
	}
}
