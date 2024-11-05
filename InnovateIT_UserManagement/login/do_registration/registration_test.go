package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func init() { // 设置哈希字段
	key := "email_captcha"
	mylink.FileConfigJSON("/home/gjz/项目/go/InnovateIT/Configuration/configure.json") //初始化配置文件
	link, err := mylink.NewRedisLink(0)                                              //初始化redis
	_, err1 := mylink.NewMysqlLink()
	if err != nil || err1 != nil {
		fmt.Println("redis连接失败")
	}

	// 设置TTL
	ttl := 30 * time.Microsecond //过期时间设置为 30秒
	err = link.Client.Expire(context.Background(), key, ttl).Err()
	if err != nil {
		fmt.Println("设置TTL失败:", err)
		return
	} //初始化mysql
	fmt.Print("init")
}

func Test_Registration_email_send(t *testing.T) {
	root := tool.NewLiabilitylist(8) //发送链路
	root.AddNode(Unique_email_redis)
	root.AddNode(Unique_email_mysql)
	root.AddNode(Addcache_email)
	root.AddNode(Captcha_email_send)
	err, s, bytes := root.RunNodeList("1820284294@qq.com$test", "$")
	if err != nil {
		log.Fatalf("HGET error: %v", err)
	}

	//打印
	fmt.Println(s)
	fmt.Printf("%v\n", bytes)
}

func TestUnique_email_mysql(t *testing.T) {
	Unique_email_mysql("1820284294@qq.com$test")
}
