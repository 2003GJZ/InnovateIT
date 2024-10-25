package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
	"database/sql"
)

//链路1 TODO email$
//验证唯一性registration_captcha
//TODO 查询缓存 email_username 将username计算md5，没有为 "NULL"       --------
//					   													|
//TODO 查询数据库														|
//TODO 插入缓存 email_username 将name计算md5 没有为"NULL"					|
//																		|
//TODO 验证码验证(发送)

//链路2
//正式加添
//TODO 验证码验证(验证)
//TODO 清除缓存 email_username
//TODO 插入数据库
//
//缓存一致性保证
//TODO 再次查询数据库
//TODO 刷新缓存 email_username login_email

// 查询数据库
func Unique_email_mysql(string2 string) (error, string, string, byte, bool) {
	//字符串切割
	email, remainder, err2 := tool.SplitString(string2, "$")

	if err2 != nil {
		return err2, "", "", 0, false
	}

	var username string
	query := "SELECT username FROM user_email_login WHERE email = ?"
	err := mylink.Sqldb.QueryRow(query, email).Scan(&username)
	//TODO 查询数据库
	log := "Unique_email:"
	if err == sql.ErrNoRows {
		//没找到，可以插入
		log += "成功"
		remainder = email + "$" + tool.GetMd5(username) + "$" + remainder
		return nil, log, remainder, 1, true
	}
	log += "已存在"
	remainder = email + "$" + "NULL" + "$"
	return err, log, remainder, 0, false
}

// 缓存查询验证
func Unique_email_redis(ags string) (error, string, string, byte, bool) {
	//字符串切割
	email, _, err2 := tool.SplitString(ags, "$")

	if err2 != nil {
		return err2, "", "", 0, false
	}

	link, err2 := mylink.NewredisLink(0)
	if err2 != nil {
		return err2, "", "", 0, false
	}
	var username string
	link.Client.HGet(link.Ctx, "email_username", email).Scan(&username)
	log := "Unique_email_redis:"
	if username == "NULL" {

		//不存在可以注册
		log += "ok"
		return nil, log, ags, 5, true
	} else if username == "" {
		//未找到查询数据库
		log += "not in redis, select database"
		return nil, log, ags, 1, true
	}
	log += "已存在"
	return nil, log, "", 0, false

}
