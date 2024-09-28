package mylink

import (
	"database/sql"
	"log"
)

func mysqllink() {
	//创建数据库连接字符串

}

func connectDB() (*sql.DB, error) {
	dsn := "用户名:密码@tcp(127.0.0.1:3306)/数据库名"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("连接数据库失败: %v", err)
		return nil, err
	}
	log.Println("数据库连接成功")
	return db, nil
}
