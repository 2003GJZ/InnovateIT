package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

// 2.5刷新缓存 (email_username) 邮箱$md5(用户名)or NULL$用户名$密码----->邮箱$用户名$密码
func Flushedcache_email(string2 string) (error, tool.Outcome) {
	//1.3_将mysql查到的信息加入缓存
	logs := "Flushedcache_email:"
	outcometmp := tool.Outcome{
		logs, "", 0, false,
	}
	email, s, err2 := tool.SplitString(string2, "$")
	usernameMd5, s2, err3 := tool.SplitString(s, "$")
	username, s3, err4 := tool.SplitString(s2, "$")
	password, _, err5 := tool.SplitString(s3, "$")
	if err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		outcometmp.Output = logs + "SplitStringERR"
		return err2, outcometmp
	}

	link, _ := mylink.NewredisLink(0)
	htable := tool.Redis_htable{
		"email_username",
		link,
	}
	htable.Insert_caching(email, usernameMd5) //插入缓存email_username
	//TODO 后面改成token
	if usernameMd5 == "NULL" {
		//插入失败更新缓存
		outcometmp.Goon = true
		outcometmp.Nextinput = email + "$" + username + "$" + "NULL"
		outcometmp.Output = logs + "插入失败"

	} else {
		//1.3_存在不再继续注册
		outcometmp.Goon = true
		outcometmp.Nextinput = email + "$" + username + "$" + password
		outcometmp.Output = logs + "注册成功缓存已插入"

	}
	return nil, outcometmp
}
