package domain

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/msdb"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/dao"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/model"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/repo"
)

type MemberAddressDomain struct {
	memberAddressRepo repo.MemberAddressRepo
}

func (d *MemberAddressDomain) FindAddressList(
	ctx context.Context,
	userId int64,
	coinId int64) ([]*model.MemberAddress, error) {
	return d.memberAddressRepo.FindByMemIdAndCoinId(ctx, userId, coinId)
}

func NewMemberAddressDomain(db *msdb.MsDB) *MemberAddressDomain {
	return &MemberAddressDomain{
		memberAddressRepo: dao.NewMemberAddressDao(db),
	}
}
