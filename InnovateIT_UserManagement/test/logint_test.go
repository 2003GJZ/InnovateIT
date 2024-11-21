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
	link, err2 := mylink.NewRedisLink(0)
	if err2 != nil {
		log.Fatalf("HGET error: %v", err2)
	}

	link.Client.HSet(link.Ctx, "login_phone", "123456789", tool.GetMd5("123456789"))

}
func Test_test(t *testing.T) {

	root := tool.NewLiabilitylist(8)
	root.AddNode(do_login.Login_redis_phone)
	root.AddNode(do_login.Login_mysql_phone)
	root.AddNode(do_login.Login_updatacache_phone)
	err2, s, bytes := root.RunNodeList("4551555145$areydssok111125595$", "$")

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
	_, err = mysqllink.Exec(query, "4551555145", "test1", tool.GetMd5("areyouok111125595")) // 使用Exec来执行插入操作
	if err != nil {
		// 插入操作失败
		t.Errorf("Failed to insert data into user_login: %v", err)
	}
}
