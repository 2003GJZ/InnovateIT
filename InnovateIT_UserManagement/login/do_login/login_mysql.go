package do_login

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Login_mysql(string2 string) (error, string, string, byte, bool) {
	//字符串切割
	phone, s, err2 := tool.SplitString(string2, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}
	password, s2, err2 := tool.SplitString(s, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}
	//md5计算
	var passwordMD5 string
	query := `SELECT password FROM user_login WHERE phone = ?`

	err := mylink.Sqldb.QueryRow(query, phone).Scan(&passwordMD5)
	if err != nil {
		if err == sql.ErrNoRows {
			//没找到，缓存更新
		} else {
			log.Fatal(err)
			return err, "database_not_have", s2, 0, false
		}

	}
	passwordCompare := tool.CompareMD5(password, passwordMD5)
	var succeed byte
	fan := ""
	if passwordCompare {
		succeed = 1
		fan = "ok"
		//写给下一个,更新缓存
		s2 = phone + "$" + passwordMD5
	} else {
		s2 = phone + "$nil$" //避免缓存失效
	}
	return nil, fan, s2, succeed, true

}

//da210d95cad66a088872e6ed0199e6fc
//da210d95cad66a088872e6ed0199e6fc
