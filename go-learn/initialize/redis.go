package initialize

import (
	"fmt"
	"go-learn/global"

	"github.com/go-redis/redis"
)

// Redis Connection
func Redis() {
	redisConfig := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Errorf("Fatal redis connection: %s", err))
	} else {
		global.REDIS = client
	}
}
