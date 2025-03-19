package handler

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange-api/internal/logic"
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange-api/internal/svc"
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange-api/internal/types"
	common "github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/tools"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type OrderHandler struct {
	svcCtx *svc.ServiceContext
}

func NewOrderHandler(svcCtx *svc.ServiceContext) *OrderHandler {
	return &OrderHandler{
		svcCtx: svcCtx,
	}
}

func (h *OrderHandler) History(w http.ResponseWriter, r *http.Request) {
	var req types.ExchangeReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewOrderLogic(r.Context(), h.svcCtx)
	resp, err := l.History(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *OrderHandler) Current(w http.ResponseWriter, r *http.Request) {
	var req types.ExchangeReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewOrderLogic(r.Context(), h.svcCtx)
	resp, err := l.Current(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *OrderHandler) Add(w http.ResponseWriter, r *http.Request) {
	var req types.ExchangeReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewOrderLogic(r.Context(), h.svcCtx)
	resp, err := l.AddOrder(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
