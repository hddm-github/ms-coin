package processor

import (
	"encoding/json"
	"market-api/internal/database"
	"market-api/internal/model"
)

const KLINE1M = "kline_1m"
const KLINE = "kline"
const TRADE = "trade"

type Processor interface {
	Process(data ProcessData)
	AddHandler(h MarketHandler)
}

type MarketHandler interface {
	HandleTrade(symbol string, data []byte)
	HandleKline(symbol string, kline *model.Kline)
}
type ProcessData struct {
	Type string //trade 交易 kline k线
	Key  []byte
	Data []byte
}
type DefaultProcessor struct {
	kafkaCli *database.KafkaClient
	handlers []MarketHandler
}

func NewDefaultProcessor(kafkaCli *database.KafkaClient) *DefaultProcessor {
	return &DefaultProcessor{
		kafkaCli: kafkaCli,
		handlers: make([]MarketHandler, 0),
	}
}

func (d *DefaultProcessor) Process(data ProcessData) {
	if data.Type == KLINE {
		symbol := string(data.Key)
		kline := &model.Kline{}
		json.Unmarshal(data.Data, kline)
		for _, v := range d.handlers {
			v.HandleKline(symbol, kline)
		}
	}

}

func (p *DefaultProcessor) AddHandler(h MarketHandler) {
	// 发送到 websocket 服务
	p.handlers = append(p.handlers, h)
}

func (p *DefaultProcessor) Init() {
	p.startReadFromKafka(KLINE1M, KLINE)
}

func (p *DefaultProcessor) startReadFromKafka(topic string, tp string) {
	// 一定要先 start 后 read
	p.kafkaCli.StartRead(topic)
	go p.dealQueueData(p.kafkaCli, tp)
}

func (p *DefaultProcessor) dealQueueData(cli *database.KafkaClient, tp string) {
	// 这就是队列的数据
	for {
		msg := cli.Read()
		data := ProcessData{
			Type: tp,
			Key:  msg.Key,
			Data: msg.Data,
		}
		p.Process(data)
	}
}
