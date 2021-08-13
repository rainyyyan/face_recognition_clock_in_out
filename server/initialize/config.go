package initialize

import (
	"server/global"

	"github.com/spf13/viper"
)

func ViperRead() {
	v := viper.New()
	v.SetConfigFile("./config.toml")
	v.SetConfigType("toml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(&global.Config)
	if err != nil {
		panic(err)
	}
}
