package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
	"fmt"
	"log"
	"testing"
)

func init() {
	mylink.FileConfigJSON("/home/gjz/项目/go/InnovateIT/Configuration/configure.json") //初始化配置文件
	mylink.NewredisLink(0)                                                           //初始化redis
	mylink.NewmysqlLink()                                                            //初始化mysql
	fmt.Print("init")
}

func Test_Registration_email_send(t *testing.T) {
	root := tool.NewLiabilitylist(8) //发送链路
	root.AddNode(Unique_email_redis)
	root.AddNode(Unique_email_mysql)
	root.AddNode(Addcache_email)
	root.AddNode(Captcha_email_send)
	err, s, bytes := root.RunNodeList("2835759623@qq.com$test", "$")
	if err != nil {
		log.Fatalf("HGET error: %v", err)
	}

	//打印
	fmt.Println(s)
	fmt.Printf("%v\n", bytes)
}

func TestUnique_email_mysql(t *testing.T) {
	Unique_email_mysql("2835759623@qq.com$test")
}
