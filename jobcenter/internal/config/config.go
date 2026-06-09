// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import (
	"jobcenter/internal/database"
	"jobcenter/internal/kline"
)

type Config struct {
	Okx   kline.OkxConfig
	Mongo database.MongoConfig
}
