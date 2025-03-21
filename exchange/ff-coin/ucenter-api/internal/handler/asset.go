package handler

import (
	common "github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/tools"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/logic"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/svc"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type AssetHandler struct {
	svcCtx *svc.ServiceContext
}

func NewAssetHandler(svcCtx *svc.ServiceContext) *AssetHandler {
	return &AssetHandler{
		svcCtx: svcCtx,
	}
}

func (h *AssetHandler) FindWalletBySymbol(w http.ResponseWriter, r *http.Request) {
	var req types.AssetReq
	if err := httpx.ParsePath(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewAssetLogic(r.Context(), h.svcCtx)
	resp, err := l.FindWalletBySymbol(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *AssetHandler) FindWallet(w http.ResponseWriter, r *http.Request) {
	var req = types.AssetReq{}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewAssetLogic(r.Context(), h.svcCtx)
	resp, err := l.FindWallet(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *AssetHandler) ResetAddress(w http.ResponseWriter, r *http.Request) {
	var req types.AssetReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewAssetLogic(r.Context(), h.svcCtx)
	resp, err := l.ResetAddress(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *AssetHandler) FindTransaction(w http.ResponseWriter, r *http.Request) {
	var req types.AssetReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewAssetLogic(r.Context(), h.svcCtx)
	resp, err := l.FindTransaction(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
