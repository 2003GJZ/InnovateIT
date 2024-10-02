package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"login/mylink"
	"login/tool"
)

func login_mysql(string2 string) (error, string, string, byte, bool) {

	phone, s, err2 := tool.SplitString(string2, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}
	password, s2, err2 := tool.SplitString(s, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}

	var passwordMD5 string
	query := `SELECT password FROM user_login WHERE phone = ?`
	err := mylink.Sqldb.QueryRow(query, phone).Scan(&passwordMD5)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("用户未找到")
		} else {
			log.Fatal(err)
		}
		return err, "database_not_have", s2, 0, false
	} else {
		passwordCompare := tool.CompareMD5(passwordMD5, password)
		var succeed byte
		if passwordCompare {
			succeed = 1
		}
		return nil, "ok", s2, succeed, passwordCompare
	}
}
