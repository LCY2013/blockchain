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

type MarketHandler struct {
	svcCtx *svc.ServiceContext
}

func NewMarketHandler(svcCtx *svc.ServiceContext) *MarketHandler {
	return &MarketHandler{
		svcCtx: svcCtx,
	}
}

func (h *MarketHandler) SymbolThumbTrend(w http.ResponseWriter, r *http.Request) {
	var req = &types.MarketReq{}
	newResult := common.NewResult()
	//获取一下ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.SymbolThumbTrend(req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *MarketHandler) SymbolThumb(w http.ResponseWriter, r *http.Request) {
	var req = &types.MarketReq{}
	newResult := common.NewResult()
	//获取一下ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.SymbolThumb(req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *MarketHandler) SymbolInfo(w http.ResponseWriter, r *http.Request) {
	var req types.MarketReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	newResult := common.NewResult()
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.SymbolInfo(req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *MarketHandler) CoinInfo(w http.ResponseWriter, r *http.Request) {
	var req types.MarketReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.CoinInfo(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *MarketHandler) History(w http.ResponseWriter, r *http.Request) {
	var req types.MarketReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.History(&req)
	result := common.NewResult().Deal(resp.List, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
