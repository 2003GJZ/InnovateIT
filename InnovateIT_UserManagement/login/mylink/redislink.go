package mylink

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"io/ioutil"
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

// 读取配置文件
func FileConfigJSON(configurl string) (error, Config) {
	// 读取 JSON 配置文件
	jsonData, err := ioutil.ReadFile(configurl)
	if err != nil {
		log.Fatalf("读取文件错误: %v", err)
	}

	var config Config
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		log.Fatalf("解析 JSON 错误: %v", err)
	}

	// 打印解析得到的配置信息
	log.Printf("Redis 主机: %s, 端口: %d, 密码: %s\n",
		config.Redis.Host, config.Redis.Port, config.Redis.Password)
	log.Printf("MySQL 主机: %s, 端口: %d, 用户: %s, 密码: %s, 数据库: %s\n",
		config.MySQL.Host, config.MySQL.Port, config.MySQL.User, config.MySQL.Password, config.MySQL.Database)

	return err, config
}

type RedisLink struct {
	Ctx    context.Context
	Client *redis.Client
}

func NewredisLink(configurl string) (error, *RedisLink) {
	err, config := FileConfigJSON(configurl)
	if err != nil {
		return err, nil
	}
	// 创建 Redis 客户端
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       0, // 默认数据库为 0
	})
	// 测试链接
	if !just_once {
		_, err = rdb.Ping(ctx).Result()
		if err != nil {
			return err, nil
		}
		log.Println("Redis 链接成功")
		just_once = true
	}

	//创建链接不测试链接
	redislink := &RedisLink{
		Ctx:    ctx,
		Client: rdb,
	}

	return nil, redislink
}

func (redislink *RedisLink) Redis_hset_get() {
	//TODO
}
