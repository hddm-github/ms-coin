// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"grpc-common/market/types/market"
	"market-api/internal/svc"
	"market-api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type MarketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (l *MarketLogic) SymbolThumbTrend(req *types.MarketReq) (list []*types.CoinThumbResp, err error) {
	symbolThumbTrend, err := l.svcCtx.MarketRpc.FindSymbolThumbTrend(context.Background(),
		&market.MarketReq{
			Ip: req.Ip,
		})
	if err != nil {
		return nil, err
	}
	if err := copier.Copy(&list, symbolThumbTrend.List); err != nil {
		return nil, err
	}
	return
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
