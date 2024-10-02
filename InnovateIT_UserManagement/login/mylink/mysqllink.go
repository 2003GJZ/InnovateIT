package mylink

import (
	"database/sql"
	"fmt"
	"log"
)

func Newmysqllink() (*sql.DB, error) {
	if Sqldb == nil {
		if configure == nil {

			return nil, fmt.Errorf("配置文件未加载")
		} else {
			dsn := configure.MySQL.User + ":" + configure.MySQL.Password + "@tcp(" + string(configure.MySQL.Host) + ":" + string(configure.MySQL.Port) + ")/" + configure.MySQL.Database
			db, err := sql.Open("mysql", dsn)
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
			Sqldb = db

		}

	}
	return Sqldb, nil
}

func GetSqldb() (*sql.DB, error) {
	if Sqldb == nil {

		db, err := Newmysqllink() //初始化数据库连接
		if err != nil {
			//数据库链接失败
			return nil, err
		}
		Sqldb = db
	}
	return Sqldb, nil
}

//func connectDB() (*sql.DB, error) {
//	dsn := "用户名:密码@tcp(127.0.0.1:3306)/数据库名"
//	db, err := sql.Open("mysql", dsn)
//	if err != nil {
//		log.Printf("连接数据库失败: %v", err)
//		return nil, err
//	}
//	log.Println("数据库连接成功")
//	return db, nil
//}
