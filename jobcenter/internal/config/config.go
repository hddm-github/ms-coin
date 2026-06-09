// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

type Config struct {
	Okx OkxConfig
}

type OkxConfig struct {
	ApiKey    string
	SecretKey string
	Pass      string
	Host      string
	Proxy     string
}
