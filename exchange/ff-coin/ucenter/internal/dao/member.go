package dao

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/msdb"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/msdb/gorms"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/model"
	"gorm.io/gorm"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func (m *MemberDao) FindMemberById(ctx context.Context, memberId int64) (mem *model.Member, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.Member{}).Where("id=?", memberId).Take(&mem).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func (m *MemberDao) UpdateLoginCount(ctx context.Context, id int64, step int) error {
	session := m.conn.Session(ctx)
	//login_count = login_count+1
	err := session.Exec("update member set login_count = login_count+? where id=?", step, id).Error
	return err
}

func (m *MemberDao) Save(ctx context.Context, mem *model.Member) error {
	session := m.conn.Session(ctx)
	err := session.Save(mem).Error
	return err
}

func (m *MemberDao) FindByPhone(ctx context.Context, phone string) (mem *model.Member, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.Member{}).
		Where("mobile_phone=?", phone).Limit(1).
		Take(&mem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return mem, err
}

func NewMemberDao(db *msdb.MsDB) *MemberDao {
	return &MemberDao{
		conn: gorms.New(db.Conn),
	}
}
