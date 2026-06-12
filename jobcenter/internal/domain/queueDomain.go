package domain

import (
	"encoding/json"
	"jobcenter/internal/database"
	"jobcenter/internal/model"
	"log"
)

const KLINE1M = "kline_1m"

type QueueDomain struct {
	kafkaCli *database.KafkaClient
}

func NewQueueDomain(kafkaCli *database.KafkaClient) *QueueDomain {
	return &QueueDomain{
		kafkaCli: kafkaCli,
	}
}

func (d *QueueDomain) Send1mKline(data []string, symbol string) {
	kline := model.NewKline(data, "1m")
	bytes, _ := json.Marshal(kline)
	msg := database.KafkaData{
		Topic: KLINE1M,
		Data:  bytes,
		Key:   []byte(symbol),
	}
	d.kafkaCli.Send(msg)
	log.Println("Kafka发送数据成功")
}
