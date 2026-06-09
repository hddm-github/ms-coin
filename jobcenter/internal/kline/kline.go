package kline

import (
	"encoding/json"
	"jobcenter/internal/config"
	"log"
	"mscoin-common/tools"
	"sync"
	"time"
)

type Kline struct {
	wg sync.WaitGroup
	c  config.OkxConfig
}
type OkxResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

func (k *Kline) Do(period string) {
	k.wg.Add(2)
	// 获取某个币 BTC-USDT ETH-USDT
	go k.getKlineData("BTC-USDT", period)
	go k.getKlineData("ETH-USDT", period)
}

func (k *Kline) getKlineData(instId string, period string) {
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

	log.Println("=================获取到的 K 线数据===============")
	log.Println("instId:", instId, "period:", period)
	log.Println("result kline data:", string(resp))
	log.Println("=================END===============")
	return
}

func NewKline(c config.OkxConfig) *Kline {
	return &Kline{
		c: c,
	}
}
