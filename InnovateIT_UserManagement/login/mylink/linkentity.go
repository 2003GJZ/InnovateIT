package mylink

var redis_just_once bool //redis 只连接测试，仅仅一次
var configure *Config    //配置文件

// MySQL 连接
type MySQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

type Config struct {
	MySQL `json:"mySQL"`
	Redis `json:"redis"`
}
