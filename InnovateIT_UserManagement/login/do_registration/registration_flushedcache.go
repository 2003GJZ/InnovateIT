package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

// 2.5刷新缓存 email_username 邮箱$md5(用户名)or NULL----->邮箱$XXX
func Flushedcache_email(string2 string) (error, tool.Outcome) {
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
	log := "Addcache_email:"
	link, _ := mylink.NewredisLink(0)
	link.Client.HSet(link.Ctx, "email_username", email, usernameMd5)
	log += "ok"
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
