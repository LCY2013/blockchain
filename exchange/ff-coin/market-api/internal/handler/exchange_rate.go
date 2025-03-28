package handler

import (
	common "github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/tools"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market-api/internal/logic"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market-api/internal/svc"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type ExchangeRateHandler struct {
	svcCtx *svc.ServiceContext
}

func NewExchangeRateHandler(svcCtx *svc.ServiceContext) *ExchangeRateHandler {
	return &ExchangeRateHandler{
		svcCtx: svcCtx,
	}
}

func (h *ExchangeRateHandler) UsdRate(w http.ResponseWriter, r *http.Request) {
	var req types.RateRequest
	if err := httpx.ParsePath(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	newResult := common.NewResult()
	//获取一下ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewExchangeRateLogic(r.Context(), h.svcCtx)
	resp, err := l.UsdRate(&req)
	result := newResult.Deal(resp.Rate, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
