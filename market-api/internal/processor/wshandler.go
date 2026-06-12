package processor

import (
	"market-api/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebSocketHandler struct {
}

func (w WebSocketHandler) HandleTrade(symbol string, data []byte) {
	//TODO implement me
	panic("implement me")
}

func (w WebSocketHandler) HandleKline(symbol string, kline *model.Kline) {
	//TODO implement me
	logx.Info("===============WebSocketHandler Start===============")
	logx.Info("symbol: ", symbol)
	logx.Info("close:", kline.ClosePrice, "high:", kline.HighestPrice, "time:", kline.Time)
	logx.Info("===============WebSocketHandler End===============")

}

func NewWebSocketHandler() *WebSocketHandler {
	return &WebSocketHandler{}
}
