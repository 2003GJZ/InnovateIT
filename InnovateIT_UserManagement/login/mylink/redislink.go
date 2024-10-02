package mylink

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

/*
if err == redis.Nil {
		fmt.Println("field3 does not exist")
	} else if err != nil {
		log.Fatalf("HGET error: %v", err)
	} else {
		fmt.Println("field1:", val1)
	}
*/
// 定义结构体以匹配 JSON 文件的结构

type RedisLink struct {
	Ctx    context.Context
	Client *redis.Client
}

func NewredisLink(databasenum int) (*RedisLink, error) {
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
	if !redis_just_once {
		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			return nil, err
		}
		log.Println("Redis 链接成功")
		redis_just_once = true
	}

	//创建链接不测试链接
	redislink := &RedisLink{
		Ctx:    ctx,
		Client: rdb,
	}

	return redislink, nil
}
