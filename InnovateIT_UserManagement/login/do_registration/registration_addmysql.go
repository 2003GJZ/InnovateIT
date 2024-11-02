package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

// 2.3插入数据库    邮箱$用户名$密码----->邮箱$用户名$密码
func Addmysql_email(string2 string) (error, tool.Outcome) {

	logs := "Addmysql_email:"
	outcometmp := tool.Outcome{
		logs, "", 0, false,
	}
	email, s, err2 := tool.SplitString(string2, "$")
	password, s2, err3 := tool.SplitString(s, "$")
	user, _, err4 := tool.SplitString(s2, "$")
	if err2 != nil || err3 != nil || err4 != nil {
		outcometmp.Output = logs + "SplitStringERR"
		return err2, outcometmp
	}

	//md5计算
	passwordMD5 := tool.GetMd5(password)

	//新增到user_email_login表
	query := "INSERT INTO user_email_login (email,password,username) VALUES (?,?,?)"
	err := mylink.Sqldb.QueryRow(query, email, passwordMD5, user)
	if err != nil {
		//注册失败，刷新缓存
		outcometmp.Output = logs + "注册失败"
		outcometmp.Bitmap = 0

	} else {
		outcometmp.Output = logs + "注册成功"
		outcometmp.Bitmap = 1
	}
	outcometmp.Goon = true
	outcometmp.Nextinput = string2

	return nil, outcometmp

}
