package initialize

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"go-web/global"
)

func InitConfig() {
	configFileName := "D:\\goproject\\go-web\\config\\config.yaml"

	v := viper.New()
	v.SetConfigFile(configFileName)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//fmt.Println(v)
	if err := v.UnmarshalKey("nacos", global.NacosConfig); err != nil {
		panic(err)
	}
	log.Info(global.NacosConfig)

	if err := v.UnmarshalKey("database", global.DatabaseConfig); err != nil {
		panic(err)
	}
	log.Info(global.DatabaseConfig)
}
