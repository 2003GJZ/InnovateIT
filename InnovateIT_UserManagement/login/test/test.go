package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-redis/redis/v8"
)

// 定义结构体以匹配 JSON 文件的结构
type Config struct {
	Redis struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Password string `json:"password"`
	} `json:"redis"`
	MySQL struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"mysql"`
}

func main() {
	// 读取 JSON 配置文件
	jsonData, err := ioutil.ReadFile("/home/gjz/项目/go/InnovateIT/Configuration/configure.json")
	if err != nil {
		log.Fatalf("读取文件错误: %v", err)
	}

	var config Config
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		log.Fatalf("解析 JSON 错误: %v", err)
	}

	// 打印解析得到的配置信息
	fmt.Printf("Redis 主机: %s, 端口: %d, 密码: %s\n", config.Redis.Host, config.Redis.Port, config.Redis.Password)
	fmt.Printf("MySQL 主机: %s, 端口: %d, 用户: %s, 密码: %s, 数据库: %s\n",
		config.MySQL.Host, config.MySQL.Port, config.MySQL.User, config.MySQL.Password, config.MySQL.Database)

	// 连接到 Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       0, // 默认数据库为 0
	})

	// 测试连接
	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("连接到 Redis 错误: %v", err)
	}
	fmt.Println("已连接到 Redis:", pong)
}
