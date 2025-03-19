package handler

import (
	common "github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/logic"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/svc"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type ApproveHandler struct {
	svcCtx *svc.ServiceContext
}

func (h *ApproveHandler) SecuritySetting(w http.ResponseWriter, r *http.Request) {
	var req types.ApproveReq
	l := logic.NewApproveLogic(r.Context(), h.svcCtx)
	resp, err := l.FindSecuritySetting(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func NewApproveHandler(svcCtx *svc.ServiceContext) *ApproveHandler {
	return &ApproveHandler{svcCtx}
}
