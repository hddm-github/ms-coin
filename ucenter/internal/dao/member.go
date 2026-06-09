package dao

import (
	"context"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
	"ucenter/internal/model"

	"gorm.io/gorm"
)

type MemberRepo struct {
	conn *gorms.GormConn
}

func (m *MemberRepo) UpdateLoginCount(ctx context.Context, id int64, step int) error {
	session := m.conn.Session(ctx)
	err := session.Exec("update member set login_count = login_count + ? where id = ?", step, id).Error
	if err != nil {
		return err
	}
	return nil
}

func NewMemberRepo(db *msdb.MsDB) *MemberRepo {
	return &MemberRepo{
		conn: gorms.New(db.Conn),
	}
}

func (m *MemberRepo) FindByPhone(ctx context.Context, phone string) (mem *model.Member, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(model.Member{}).
		Where("mobile_phone = ?", phone).
		Limit(1).
		Take(&mem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return mem, nil
}

func (m *MemberRepo) Save(ctx context.Context, mem *model.Member) (err error) {
	session := m.conn.Session(ctx)
	err = session.Save(mem).Error
	return err
}
