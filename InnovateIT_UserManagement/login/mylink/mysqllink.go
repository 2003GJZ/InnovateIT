package mylink

import (
	"database/sql"
	"log"
)

type Mysqllink struct {
	Sqldb *sql.DB
}

func Newmysqllink() *sql.DB {
	if configure == nil {
		log.Fatal("配置文件未加载")
		return nil
	} else {
		dsn := configure.MySQL.User + ":" + configure.MySQL.Password + "@tcp(" + string(configure.MySQL.Host) + ":" + string(configure.MySQL.Port) + ")/" + configure.MySQL.Database
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return db
	}

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
