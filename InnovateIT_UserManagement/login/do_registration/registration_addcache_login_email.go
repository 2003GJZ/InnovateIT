package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

// 2.6将缓存写入 login_email 邮箱$用户名$密码---->end
func Addcache_login_email(string2 string) (error, tool.Outcome) {
	logs := "Flushedcache_login_email:"
	outcometmp := tool.Outcome{
		logs, "", 0, false,
	}
	email, s, err2 := tool.SplitString(string2, "$")
	_, s2, err3 := tool.SplitString(s, "$")
	password, s3, err4 := tool.SplitString(s2, "$")
	if err2 != nil || err3 != nil || err4 != nil {
		outcometmp.Output = logs + "SplitStringERR"
		return err2, outcometmp
	}
	outcometmp.Nextinput = s3

	link, _ := mylink.NewredisLink(0)
	htable := tool.Redis_htable{
		"login_email",
		link,
	}
	//md5计算
	if password != "NULL" {
		password = tool.GetMd5(password)
	}

	htable.Insert_caching(email, password)
	return nil, outcometmp

}
