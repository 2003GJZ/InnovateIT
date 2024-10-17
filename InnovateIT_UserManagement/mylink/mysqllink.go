package mylink

import (
	"database/sql"
	"fmt"
)

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
