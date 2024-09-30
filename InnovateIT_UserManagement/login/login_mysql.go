package main

import (
	_ "github.com/go-sql-driver/mysql"
)

// phone $ password
// xxxx$hjjhjjh$
//func login_mysql(string2 string) (error, string, string, byte, bool) { //查redis
//
//	phone, s, err2 := tool.SplitString(string2, "$")
//	if err2 != nil {
//		return err2, "", "", 0, false
//	}
//	password, s2, err2 := tool.SplitString(s, "$")
//	if err2 != nil {
//		return err2, "", "", 0, false
//	}
//
//}

//	func main() {
//		// 创建数据库连接字符串
//		myapp := sqlurl{name: "root", password: "123456", host: "localhost", port: "3306"}
//
//		// 连接数据库
//		db, err := sql.Open("mysql", myapp.tosqlurl("myapp"))
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer db.Close()
//
//		// 测试数据库连接
//		err = db.Ping()
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		// 从用户输入获取登录信息
//		fmt.Print("Enter phone number: ")
//		var phoneNumber string
//		fmt.Scanln(&phoneNumber)
//		fmt.Print("Enter password: ")
//		var password string
//		fmt.Scanln(&password)
//
//		// 准备SQL查询
//		var username string
//		query := `SELECT username FROM user_login_phone WHERE phone_number=? AND password=?`
//
//		// 执行查询
//		err = db.QueryRow(query, phoneNumber, password).Scan(&username)
//
//		if err != nil {
//
//			if err == sql.ErrNoRows {
//				fmt.Println("Login failed: User not found or incorrect password.")
//			} else {
//				log.Fatal(err)
//			}
//		} else {
//			fmt.Printf("Login successful for user: %s\n", username)
//		}
//	}
