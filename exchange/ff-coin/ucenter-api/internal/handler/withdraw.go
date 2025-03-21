package handler

import (
	common "github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/logic"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/svc"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type WithdrawHandler struct {
	svcCtx *svc.ServiceContext
}

func (h *WithdrawHandler) QueryWithdrawCoin(w http.ResponseWriter, r *http.Request) {
	var req types.WithdrawReq
	l := logic.NewWithdrawLogic(r.Context(), h.svcCtx)
	resp, err := l.QueryWithdrawCoin(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *WithdrawHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	var req types.WithdrawReq
	l := logic.NewWithdrawLogic(r.Context(), h.svcCtx)
	resp, err := l.SendCode(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *WithdrawHandler) WithdrawCode(w http.ResponseWriter, r *http.Request) {
	var req types.WithdrawReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	l := logic.NewWithdrawLogic(r.Context(), h.svcCtx)
	resp, err := l.WithdrawCode(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *WithdrawHandler) Record(w http.ResponseWriter, r *http.Request) {
	var req types.WithdrawReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	l := logic.NewWithdrawLogic(r.Context(), h.svcCtx)
	resp, err := l.Record(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func NewWithdrawHandler(svcCtx *svc.ServiceContext) *WithdrawHandler {
	return &WithdrawHandler{svcCtx}
}
