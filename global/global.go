package global

import "go-web/config"

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	DatabaseConfig = &config.MysqlConfig{}
)
