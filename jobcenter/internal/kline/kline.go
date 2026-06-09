package kline

import (
	"encoding/json"
	"jobcenter/internal/database"
	"jobcenter/internal/domain"
	"log"
	"mscoin-common/tools"
	"sync"
	"time"
)

type OkxConfig struct {
	ApiKey    string
	SecretKey string
	Pass      string
	Host      string
	Proxy     string
}

type Kline struct {
	wg          sync.WaitGroup
	c           OkxConfig
	klineDomain *domain.KlineDomain
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
	}
	log.Println("=================END===============")
	k.wg.Done()
	return
}

func NewKline(c OkxConfig, client *database.MongoClient) *Kline {
	return &Kline{
		c:           c,
		klineDomain: domain.NewklineDomain(client),
	}
}
