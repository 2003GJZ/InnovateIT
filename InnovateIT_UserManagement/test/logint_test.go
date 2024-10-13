package test

import (
	"InnovateIT_UserManagement/login/do_login"
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
	"fmt"
	"log"
	"testing"
)

func TestRedis(t *testing.T) {
	link, err2 := mylink.NewredisLink(0)
	if err2 != nil {
		log.Fatalf("HGET error: %v", err2)
	}

	link.Client.HSet(link.Ctx, "login_phone", "123456789", tool.GetMd5("123456789"))

}
func Test_test(t *testing.T) {
	mylink.FileConfigJSON("/home/gjz/项目/go/InnovateIT1/Configuration/configure.json") //初始化配置文件
	mylink.NewredisLink(0)                                                            //初始化redis
	mylink.NewmysqlLink()                                                             //初始化mysql
	fmt.Print("init")
	root := tool.NewLiabilitylist(8)
	root.AddNode(do_login.Login_redis)
	root.AddNode(do_login.Login_mysql)
	root.AddNode(do_login.Updatacache)
	err2, s, bytes := root.RunNodeList("987654321$areyouok1111255595$", "$")

	if err2 != nil {
		log.Fatalf("HGET error: %v", err2)
	}

	//打印
	fmt.Println(s)
	fmt.Printf("%v\n", bytes)
}

func TestMysql(t *testing.T) {
	mysqllink, err := mylink.GetSqldb() // 假设这是创建数据库连接的正确方法
	if err != nil {
		// MySQL链接失败
		t.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer mysqllink.Close() // 确保在函数结束时关闭数据库连接

	query := "INSERT INTO user_login(phone, username, password) VALUES (?, ?, ?)"
	_, err = mysqllink.Exec(query, "987654321", "test", tool.GetMd5("areyouok1111255595")) // 使用Exec来执行插入操作
	if err != nil {
		// 插入操作失败
		t.Errorf("Failed to insert data into user_login: %v", err)
	}
}
