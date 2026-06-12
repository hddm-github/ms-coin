// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"grpc-common/market/mclient"
	"market-api/internal/config"
	"market-api/internal/database"
	"market-api/internal/processor"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	ExchangeRateRpc mclient.ExchangeRate
	MarketRpc       mclient.Market
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化 Processor
	kafkaClient := database.NewKafkaClient(c.Kafka)
	defaultProcessor := processor.NewDefaultProcessor(kafkaClient)
	defaultProcessor.Init()
	defaultProcessor.AddHandler(processor.NewWebSocketHandler())
	return &ServiceContext{
		Config:          c,
		ExchangeRateRpc: mclient.NewRate(zrpc.MustNewClient(c.MarketRpc)),
		MarketRpc:       mclient.NewMarket(zrpc.MustNewClient(c.MarketRpc)),
	}
}
