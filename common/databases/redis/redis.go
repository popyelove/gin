package redis

import (
	"github.com/go-redis/redis"
	"log"
	"Gin/common/tools/setting"
)
var RedisClient *redis.Client
func init()  {
	var err error
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password,
		DB:       setting.RedisSetting.Db,
	})
	_,err = RedisClient.Ping().Result()
	if err!=nil{
		log.Fatalln(err)
	}
}
