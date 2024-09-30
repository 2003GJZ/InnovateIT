package mylink

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// 读取配置文件
func FileConfigJSON(configurl string) (error, *Config) {
	// 读取 JSON 配置文件
	jsonData, err := ioutil.ReadFile(configurl)
	if err != nil {
		log.Fatalf("读取文件错误: %v", err)
	}

	err = json.Unmarshal(jsonData, &configure)
	if err != nil {
		log.Fatalf("解析 JSON 错误: %v", err)
	}

	// 打印解析得到的配置信息
	log.Printf("Redis 主机: %s, 端口: %d, 密码: %s\n",
		configure.Redis.Host, configure.Redis.Port, configure.Redis.Password)
	log.Printf("MySQL 主机: %s, 端口: %d, 用户: %s, 密码: %s, 数据库: %s\n",
		configure.MySQL.Host, configure.MySQL.Port, configure.MySQL.User, configure.MySQL.Password, configure.MySQL.Database)

	return err, configure
}
