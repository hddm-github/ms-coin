package config

import (
	"jobcenter/internal/database"
)

type Config struct {
	Okx   OkxConfig
	Mongo database.MongoConfig
	Kafka database.KafkaConfig
}

type OkxConfig struct {
	ApiKey    string
	SecretKey string
	Pass      string
	Host      string
	Proxy     string
}
