package global

import (
	"go-learn/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// 全局配置
var (
	REDIS  *redis.Client
	CONFIG config.Server
	VP     *viper.Viper
	DB     *gorm.DB
)
