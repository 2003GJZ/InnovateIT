package main

import (
	"InnovateIT_UserManagement/mylink"
	"fmt"
)

func init() {

	mylink.FileConfigJSON("/home/gjz/项目/go/InnovateIT1/Configuration/configure.json") //初始化配置文件
	mylink.NewredisLink(0)                                                            //初始化redis
	mylink.NewmysqlLink()                                                             //初始化mysql
	fmt.Print("init")
}
func main() {

}
