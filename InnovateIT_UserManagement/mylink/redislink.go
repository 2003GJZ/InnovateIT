package mylink

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisLink struct {
	Ctx    context.Context
	Client *redis.Client
}

func NewRedisLink(databasenum int) (*RedisLink, error) {
	if configure == nil {
		return nil, fmt.Errorf("configure is nil")
	}
	// 创建 Redis 客户端
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", configure.Redis.Host, configure.Redis.Port),
		Password: configure.Redis.Password,
		DB:       databasenum, // 默认数据库为 0
	})
	// 测试链接
	if !REDIS_JUST_ONCE {
		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			return nil, err
		}
		log.Println("Redis 链接成功")
		REDIS_JUST_ONCE = true
	}

	//创建链接不测试链接
	redislink := &RedisLink{
		Ctx:    ctx,
		Client: rdb,
	}

	return redislink, nil
}

func GetredisLink() (*RedisLink, error) {
	return NewRedisLink(0)
}
