package handler

import (
	"market-api/internal/logic"
	"market-api/internal/svc"
	"market-api/internal/types"
	common "mscoin-common"
	"mscoin-common/tools"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type MarketHandler struct {
	svcCtx *svc.ServiceContext
}

func NewMarketHandler(svcCtx *svc.ServiceContext) *MarketHandler {
	return &MarketHandler{
		svcCtx: svcCtx}
}

func (h *MarketHandler) SymbolThumbTrend(w http.ResponseWriter, r *http.Request) {
	var req types.MarketReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	newResult := common.NewResult()
	// 获取一下 IP
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.SymbolThumbTrend(&req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
