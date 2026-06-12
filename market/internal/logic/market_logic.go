package logic

import (
	"context"
	"grpc-common/market/types/market"
	"market/internal/domain"
	"market/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	exchangeCoinDomain *domain.ExchangeCoinDomain
	marketDomain       *domain.MarketDomain
}

func (l *MarketLogic) FindSymbolThumbTrend(req *market.MarketReq) (*market.SymbolThumbRes, error) {
	coins := l.exchangeCoinDomain.FindVisible(l.ctx)
	// 查询 mongo 中对应的数据
	// 查询 1H 间隔的 可以根据时间进行查询 当天的价格趋势
	trends := l.marketDomain.FindSymbolThumbTrend(coins)
	//coinThrmbs := make([]*market.CoinThumb, len(coins))
	//for i, v := range coins {
	//	ct := &market.CoinThumb{}
	//	ct.Symbol = v.Symbol
	//	trend := make([]float64, 0)
	//	for _ = range 20 {
	//		trend = append(trend, rand.Float64())
	//	}
	//	ct.Trend = trend
	//	coinThrmbs[i] = ct
	//}
	return &market.SymbolThumbRes{
		List: trends,
	}, nil
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		exchangeCoinDomain: domain.NewExchangeCoinDomain(svcCtx.Db),
		marketDomain:       domain.NewMarketDomain(svcCtx.MongoClient),
	}
}
