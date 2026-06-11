package domain

import (
	"context"
	"market/internal/dao"
	"market/internal/model"
	"market/internal/repo"
	"mscoin-common/msdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExchangeCoinDomain struct {
	exchangeCoinRepo repo.ExchangeCoinRepo
}

func NewExchangeCoinDomain(db *msdb.MsDB) *ExchangeCoinDomain {
	return &ExchangeCoinDomain{
		exchangeCoinRepo: dao.NewExchangeCoinDao(db),
	}
}

func (d *ExchangeCoinDomain) FindVisible(ctx context.Context) (list []*model.ExchangeCoin) {
	list, err := d.exchangeCoinRepo.FindVisible(ctx)
	if err != nil {
		logx.Error(err)
		return []*model.ExchangeCoin{}
	}
	return list
}
