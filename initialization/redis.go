package initialization

import (
	"github.com/go-redis/redis"
	"github.com/wiiCoder/go-scaffold/global"
)

func Redis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.RedisConfig.Host + ":" + string(global.Config.RedisConfig.Post),
		Password: global.Config.RedisConfig.Password, // no password set
		DB:       global.Config.RedisConfig.DB,       // 默认使用 DB0
	})

	global.Rd = rdb
}
