// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"grpc-common/market/types/rate"
	"market-api/internal/svc"
	"market-api/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExchangeRateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (l *ExchangeRateLogic) UsdRate(req *types.RateRequest) (*types.RateResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 30*time.Second)
	defer cancel()

	rateRes, err := l.svcCtx.ExchangeRateRpc.UsdRate(ctx, &rate.RateReq{
		Unit: req.Unit,
		Ip:   req.Ip,
	})
	if err != nil {
		return nil, err
	}

	return &types.RateResponse{
		Rate: rateRes.Rate,
	}, nil
}

func NewExchangeRateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeRateLogic {
	return &ExchangeRateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
