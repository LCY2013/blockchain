package logic

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/tools"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/ucenter/types/login"
	"github.com/jinzhu/copier"
	"time"

	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/svc"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	loginReq := &login.LoginReq{}
	if err := copier.Copy(loginReq, req); err != nil {
		return nil, err
	}
	loginResp, err := l.svcCtx.UCLoginRpc.Login(ctx, loginReq)
	if err != nil {
		return nil, err
	}
	resp = &types.LoginRes{}
	if err := copier.Copy(resp, loginResp); err != nil {
		return nil, err
	}
	return
}

func (l *LoginLogic) CheckLogin(token string) (bool, error) {
	_, err := tools.ParseToken(token, l.svcCtx.Config.JWT.AccessSecret)
	if err != nil {
		logx.Error(err)
		return false, nil
	}
	return true, nil
}
