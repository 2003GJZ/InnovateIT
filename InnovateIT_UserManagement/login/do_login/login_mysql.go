package do_login

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Login_mysql_phone(string2 string) (error, tool.Outcome) {
	//字符串切割
	logs := "Login_mysql_phone:"
	outcometmp := tool.Outcome{
		"", "", 0, false,
	}
	phone, s, err2 := tool.SplitString(string2, "$")

	if err2 != nil {
		logs += "SplitStringERR"
		outcometmp.Output = logs
		return err2, outcometmp
	}
	password, s2, err2 := tool.SplitString(s, "$")
	if err2 != nil {
		logs += "SplitStringERR"
		outcometmp.Output = logs
		return err2, outcometmp
	}

	//md5计算
	var passwordMD5 string
	query := `SELECT password FROM user_login WHERE phone = ?`

	err := mylink.Sqldb.QueryRow(query, phone).Scan(&passwordMD5)
	if err != nil {
		if err == sql.ErrNoRows {
			//没找到，缓存更新
		} else {
			logs += "database_not_have"
			log.Fatal(err)
			outcometmp.Output = logs
			return err, outcometmp
		}

	}
	passwordCompare := tool.CompareMD5(password, passwordMD5)
	var succeed byte

	if passwordCompare {
		succeed = 1
		logs += "ok"
		//写给下一个,更新缓存
		s2 = phone + "$" + passwordMD5 + "$"
	} else {
		logs += "nil"
		s2 = phone + "$nil$" //避免缓存失效
	}
	outcometmp.Output = logs
	outcometmp.Bitmap = succeed
	outcometmp.Goon = true
	outcometmp.Nextinput = s2
	return nil, outcometmp

}

//da210d95cad66a088872e6ed0199e6fc
//da210d95cad66a088872e6ed0199e6fc
