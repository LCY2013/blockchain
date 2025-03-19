package handler

import (
	"errors"
	common "github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/tools"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/logic"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/svc"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type RegisterHandler struct {
	svcCtx *svc.ServiceContext
}

func NewRegisterHandler(svcCtx *svc.ServiceContext) *RegisterHandler {
	return &RegisterHandler{
		svcCtx: svcCtx,
	}
}

func (h *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req types.Request
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	newResult := common.NewResult()
	if req.Captcha == nil {
		httpx.OkJsonCtx(r.Context(), w, newResult.Deal(nil, errors.New("人机校验不通过")))
		return
	}
	//获取一下ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, err := l.Register(&req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *RegisterHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	var req types.CodeRequest
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, err := l.SendCode(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
