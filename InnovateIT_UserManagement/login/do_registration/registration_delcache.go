package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

// 2.2清除缓存 email_username   邮箱$用户名$密码----->邮箱$用户名$密码
func Delcache_email(string2 string) (error, tool.Outcome) {
	logs := "Delcache_email:"
	outcometmp := tool.Outcome{
		logs, "", 0, false,
	}
	email, _, err2 := tool.SplitString(string2, "$")
	if err2 != nil {
		outcometmp.Output = logs + "SplitStringERR"
		return err2, outcometmp
	}
	link, err3 := mylink.NewredisLink(0)
	if err3 != nil {
		outcometmp.Output = logs + " redis Link err"
		return err3, outcometmp
	}
	htable := tool.Redis_htable{
		Htabname:   "email_username ",
		Redis_link: link,
	}
	outcometmp.Output = logs + "ok"
	outcometmp.Bitmap = 1
	outcometmp.Goon = true
	outcometmp.Nextinput = string2
	htable.Delete_caching(email)
	return nil, outcometmp
}
