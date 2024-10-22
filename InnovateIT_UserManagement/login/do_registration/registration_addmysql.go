package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

// 插入数据库
func Addmysql_email(string2 string) (error, string, string, byte, bool) {
	email, s, err2 := tool.SplitString(string2, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}
	password, s2, err2 := tool.SplitString(s, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}
	user, _, err2 := tool.SplitString(s2, "$")
	//md5计算
	passwordMD5 := tool.GetMd5(password)
	//新增到user_email_login表
	query := "INSERT INTO user_email_login (email,password,username) VALUES (?,?,?)"
	log := "Addmysql_email:"
	err := mylink.Sqldb.QueryRow(query, email, passwordMD5, user).Scan(&passwordMD5)
	if err != nil {
		log += "注册失败"
		return err, log, "", 0, false
	}
	log += "注册成功"

	return nil, log, string2, 1, true

}
