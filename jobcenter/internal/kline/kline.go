package kline

import (
	"encoding/json"
	"jobcenter/internal/config"
	"jobcenter/internal/domain"
	"jobcenter/internal/svc"
	"log"
	"mscoin-common/tools"
	"sync"
	"time"
)

type Kline struct {
	wg          sync.WaitGroup
	c           config.OkxConfig
	klineDomain *domain.KlineDomain
	queueDomain *domain.QueueDomain
}
type OkxResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

func (k *Kline) Do(period string) {
	k.wg.Add(2)
	// 获取某个币 BTC-USDT ETH-USDT
	go k.getKlineData("BTC-USDT", "BTC/USDT", period)
	go k.getKlineData("ETH-USDT", "ETH/USDT", period)
	//go k.getKlineData("BNB-USDT", "BNB/USDT", period)
	//go k.getKlineData("SOL-USDT", "SOL/USDT", period)
	//go k.getKlineData("XRP-USDT", "XRP/USDT", period)
	//go k.getKlineData("HOGE-USDT", "HOGE/USDT", period)
	//go k.getKlineData("FIL-USDT", "FIL/USDT", period)
	//go k.getKlineData("TRX-USDT", "TRX/USDT", period)
	//go k.getKlineData("EOS-USDT", "EOS/USDT", period)
	//go k.getKlineData("ADA-USDT", "ADA/USDT", period)
}

func (k *Kline) getKlineData(instId string, symbol string, period string) {
	// 发起 http 请求  获取数据
	api := k.c.Host + "/api/v5/market/candles?instId=" + instId + "&bar" + period
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET/api/v5/market/candles?instId="+instId+"&bar"+period, k.c.SecretKey)
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = k.c.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = k.c.Pass

	resp, err := tools.GetWithHeader(api, header, k.c.Proxy)
	if err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}
	var result = &OkxResult{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}

	log.Println("=================执行存储 monogo===============")
	if result.Code == "0" {
		//代表成功
		k.klineDomain.SaveBatch(result.Data, symbol, period)
		if "1m" == period {
			// 把这个最新的数据 result.Data[0] 推送到 market 服务，推送到前端页面，进行实时变化
			if len(result.Data) > 0 {
				k.queueDomain.Send1mKline(result.Data[0], symbol)
			}
		}
	}
	log.Println("=================END===============")
	k.wg.Done()
	return
}

func NewKline(c config.OkxConfig, ctx *svc.ServiceContext) *Kline {
	return &Kline{
		c:           c,
		klineDomain: domain.NewklineDomain(ctx.MongoClient),
		queueDomain: domain.NewQueueDomain(ctx.KafkaClient),
	}
}
