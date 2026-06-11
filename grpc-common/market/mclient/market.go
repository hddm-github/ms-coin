// Code genemarketd by goctl. DO NOT EDIT.
// goctl 1.10.1
// Source: ucclient.proto

package mclient

import (
	"context"
	"grpc-common/market/types/market"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	MarketReq = market.MarketReq
	MarketRes = market.SymbolThumbRes

	Market interface {
		FindSymbolThumbTrend(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*MarketRes, error)
	}

	defaultMarket struct {
		cli zrpc.Client
	}
)

func NewMarket(cli zrpc.Client) Market {
	return &defaultMarket{
		cli: cli,
	}
}

func (m *defaultMarket) FindSymbolThumbTrend(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*MarketRes, error) {
	client := market.NewMarketClient(m.cli.Conn())
	return client.FindSymbolThumbTrend(ctx, in, opts...)
}
