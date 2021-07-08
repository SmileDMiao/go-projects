package initialize

import (
	"fmt"
	"go-learn/global"

	"github.com/spf13/viper"
)

// Viper 初始化
func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("config")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
