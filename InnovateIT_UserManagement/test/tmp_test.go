package test

import (
	"InnovateIT_UserManagement/mylink"
	"fmt"
	"testing"
)

func init() {
	mylink.FileConfigJSON("/home/gjz/项目/go/InnovateIT1/Configuration/configure.json") //初始化配置文件
	mylink.NewredisLink(0)                                                            //初始化redis
	mylink.NewmysqlLink()                                                             //初始化mysql
	fmt.Print("init")
}

func TestRedis_tmp(t *testing.T) {
	fmt.Println("sdhs")
	link, _ := mylink.NewredisLink(0)
	link.Client.HSet(link.Ctx, "test", "test", "NULL")
	var get string
	link.Client.HGet(link.Ctx, "test1", "test1").Scan(&get)
	fmt.Println("get---------------", get)
}
