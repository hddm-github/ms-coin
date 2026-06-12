package domain

import (
	"context"
	"grpc-common/market/types/market"
	"market/internal/dao"
	"market/internal/database"
	"market/internal/model"
	"market/internal/repo"
	"mscoin-common/tools"
	"sync"
	"time"
)

type MarketDomain struct {
	klineRepo repo.KlineRepo
}

func NewMarketDomain(mongoClient *database.MongoClient) *MarketDomain {
	return &MarketDomain{klineRepo: dao.NewKlineDao(mongoClient.Db)}
}

func (d *MarketDomain) FindSymbolThumbTrend(coins []*model.ExchangeCoin) []*market.CoinThumb {
	//业务模型 == rpc 传输模型
	coinThumbs := make([]*market.CoinThumb, len(coins))
	var wg sync.WaitGroup
	for i, v := range coins {
		wg.Add(1)
		go func(i int, v *model.ExchangeCoin) {
			defer wg.Done()
			from := tools.ZeroTime()
			end := time.Now().UnixMilli()
			klines, err := d.klineRepo.FindBySymbolTime(context.Background(), v.Symbol, "1H", from, end)
			if err != nil {
				coinThumbs[i] = model.DefaultCoinThumb(v.Symbol)
				return
			}
			length := len(klines)
			if length <= 0 {
				coinThumbs[i] = model.DefaultCoinThumb(v.Symbol)
				return
			}
			// 降序排列
			// 构建趋势
			trend := make([]float64, length)
			var high float64 = klines[0].HighestPrice
			var low float64 = klines[0].LowestPrice
			for j := length - 1; j >= 0; j-- {
				trend[j] = klines[j].ClosePrice
				highestPrice := klines[j].HighestPrice
				if highestPrice > high {
					high = highestPrice
				}
				lowPrice := klines[j].LowestPrice
				if lowPrice < low {
					low = lowPrice
				}
			}
			newKline := klines[0]
			oldKline := klines[length-1]
			thumb := newKline.ToCoinThumb(v.Symbol, oldKline)
			thumb.Trend = trend
			coinThumbs[i] = thumb
		}(i, v)
	}
	wg.Wait()
	return coinThumbs
}
