package mylink

import (
	"database/sql"
	"fmt"
)

//func Newmysqllink() (*sql.DB, error) {
//	// 请根据实际情况修改以下DSN
//	dsn := "gjz:20030220@tcp(127.0.0.1:3306)/user_management_db"
//	db, err := sql.Open("mysql", dsn)
//	if err != nil {
//		return nil, err
//	}
//	if err = db.Ping(); err != nil { // 检查连接
//		return nil, err
//	}
//	return db, nil
//}

func NewmysqlLink() (*sql.DB, error) {
	if Sqldb == nil {
		if configure == nil {
			return nil, fmt.Errorf("配置文件未加载")
		} else {
			// 构建 DSN
			dsn := fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s",
				configure.MySQL.User,
				configure.MySQL.Password,
				configure.MySQL.Host,
				configure.MySQL.Port,
				configure.MySQL.Database,
			)
			// 打开数据库连接
			db, err := sql.Open("mysql", dsn)
			if err != nil {
				return nil, fmt.Errorf("数据库连接失败: %v", err)
			}
			// 检查连接
			if err = db.Ping(); err != nil {
				return nil, fmt.Errorf("数据库连接测试失败: %v", err)
			}
			fmt.Println("数据库连接成功")
			Sqldb = db
		}
	}
	return Sqldb, nil
}

func GetSqldb() (*sql.DB, error) {
	if Sqldb == nil {

		db, err := NewmysqlLink() //初始化数据库连接
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
