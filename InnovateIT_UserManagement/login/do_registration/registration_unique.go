package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

//链路1  email$
//验证唯一性registration_captcha
// 1.1查询缓存 email_username 将username计算md5，没有为 "NULL"       --------
//					   													|
// 1.2查询数据库															|
// 1.3插入缓存 email_username 将name计算md5 没有为"NULL"					|
//																		|
// 1.4验证码验证(发送)

//链路2
//正式加添
// 2.1验证码验证(验证)
// 2.2清除缓存 email_username
// 2.3插入数据库
//
//缓存一致性保证
// 2.4再次查询数据库
//TODO 2.5刷新缓存 email_username
// 2.6将缓存写入 login_email

// 1.2查询数据库 邮箱$用户名---->邮箱$md5(用户名)or NULL$用户名
// 2.4再次查询数据库 邮箱$用户名$密码----->邮箱$md5(用户名)or NULL$用户名$密码
func Unique_email_mysql(string2 string) (error, tool.Outcome) {
	logs := "Unique_email_mysql:"
	outcometmp := tool.Outcome{
		logs, "", 0, false,
	}
	//字符串切割
	email, remainder, err2 := tool.SplitString(string2, "$")
	if err2 != nil {
		outcometmp.Output = logs + "SplitStringERR"
		return err2, outcometmp
	}

	var username string
	query := "SELECT username FROM user_email_login WHERE email = ?"
	mylink.Sqldb.QueryRow(query, email).Scan(&username)
	// 查询数据库
	if username == "" {
		//没找到，可以插入
		remainder = email + "$" + "NULL" + "$" + remainder
		outcometmp.Goon = true
		outcometmp.Bitmap = 1
		outcometmp.Output = logs + "成功"
		outcometmp.Nextinput = remainder
	} else {
		remainder = email + "$" + tool.GetMd5(username) + "$"
		outcometmp.Goon = false
		outcometmp.Bitmap = 0
		outcometmp.Output = logs + "已存在"
		outcometmp.Nextinput = remainder
	}

	return nil, outcometmp
}

// 1.1缓存查询验证
func Unique_email_redis(ags string) (error, tool.Outcome) {
	logs := "Unique_email_redis:"
	outcometmp := tool.Outcome{
		logs, "", 0, false,
	}
	//字符串切割
	email, _, err2 := tool.SplitString(ags, "$")

	if err2 != nil {
		outcometmp.Output = logs + "SplitStringERR"
		return err2, outcometmp
	}

	link, err3 := mylink.NewredisLink(0)
	if err2 != nil || err3 != nil {
		outcometmp.Output = logs + "Redis ERR"
		return err2, outcometmp
	}
	var username string
	link.Client.HGet(link.Ctx, "email_username", email).Scan(&username)
	if username == "NULL" {
		//不存在可以注册
		outcometmp.Output = logs + "ok"
		outcometmp.Bitmap = 5
		outcometmp.Goon = true
		outcometmp.Nextinput = ags

	} else if username == "" {
		//未找到查询数据库
		outcometmp.Output = logs + "not in redis, select database"
		outcometmp.Bitmap = 1
		outcometmp.Goon = true
		outcometmp.Nextinput = ags

	} else {
		//已存在不能注册
		outcometmp.Output = logs + "已存在"
		outcometmp.Bitmap = 0
		outcometmp.Goon = false
		outcometmp.Nextinput = ""
	}

	return nil, outcometmp

}
