package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

//registration_flushedcache.go

// 1.3加添缓存，查看是否有该用户(email_username) 邮箱$md5(用户名)or NULL$用户名$密码----->邮箱$用户名$密码

func Addcache_email(string2 string) (error, tool.Outcome) {
	//将mysql查到的信息加入缓存
	logs := "Addcache_email:"
	outcometmp := tool.Outcome{
		logs, "", 0, false,
	}
	email, s, err2 := tool.SplitString(string2, "$")
	usernameMd5, s2, err3 := tool.SplitString(s, "$")
	if err2 != nil || err3 != nil {
		outcometmp.Output = logs + "SplitStringERR"
		return err2, outcometmp
	}

	link, _ := mylink.NewredisLink(0)
	htable := tool.Redis_htable{
		"email_username",
		link,
	}
	htable.Insert_caching(email, usernameMd5) //插入缓存email_username

	if usernameMd5 == "NULL" {
		//不存在可以继续注册
		outcometmp.Goon = true
		s2 = email + "$" + s2
		outcometmp.Output = logs + "The user is not registered"
		outcometmp.Nextinput = s2
	} else {
		//存在不再继续注册
		outcometmp.Goon = false
		outcometmp.Output = logs + "The user is registered"
	}
	return nil, outcometmp
}
