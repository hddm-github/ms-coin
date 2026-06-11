package dao

import (
	"context"
	"market/internal/model"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
)

type ExchangeCoinRepo struct {
	conn *gorms.GormConn
}

func (m *ExchangeCoinRepo) FindVisible(ctx context.Context) (list []*model.ExchangeCoin, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.ExchangeCoin{}).Where("is_visible = ?", 1).Find(&list).Error
	return
}

func NewExchangeCoinRepo(db *msdb.MsDB) *ExchangeCoinRepo {
	return &ExchangeCoinRepo{
		conn: gorms.New(db.Conn),
	}
}
