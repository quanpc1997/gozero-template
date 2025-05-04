package config

import (
	"github.com/zeromicro/go-zero/core/conf"
)

type PostgresClient struct {
	Postgres struct {
		Driver          string `json:"driver"`
		Datasource      string `json:"datasource"`
		MaxIdleConns    int    `json:"maxIdleConns"`
		MaxOpenConns    int    `json:"maxOpenConns"`
		ConnMaxLifetime int    `json:"connMaxLifetime"`
	} `json:"postgres"`
}

func NewPostgresClient(filename string) (*Config, error) {
	var c Config
	err := conf.LoadConfig(filename, &c)
	return &c, err
}
