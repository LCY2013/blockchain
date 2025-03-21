package logic

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/ucenter/types/member"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/domain"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/svc"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	memberDomain *domain.MemberDomain
}

func (l *MemberLogic) FindMemberById(req *member.MemberReq) (*member.MemberInfo, error) {
	mem, err := l.memberDomain.FindMemberById(l.ctx, req.MemberId)
	if err != nil {
		return nil, err
	}
	resp := &member.MemberInfo{}
	copier.Copy(resp, mem)
	return resp, nil
}

func NewMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLogic {
	return &MemberLogic{
		ctx:          ctx,
		svcCtx:       svcCtx,
		Logger:       logx.WithContext(ctx),
		memberDomain: domain.NewMemberDomain(svcCtx.Db),
	}
}
