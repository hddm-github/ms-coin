package domain

import (
	"market/internal/repo"
	"strings"
)

type ExchangeRateDomain struct {
	exchangeCoinRepo repo.ExchangeCoinRepo
}

func NewExchangeRateDomain() *ExchangeRateDomain {
	return &ExchangeRateDomain{}
}

func (d *ExchangeRateDomain) UsdRate(unit string) float64 {
	// 应该去 redis 查询，在定时任务做一个
	unit = strings.ToUpper(unit)
	if "CNY" == unit {
		return 6.8
	} else if "JPY" == unit {
		return 180.03
	}

	return 0
}
