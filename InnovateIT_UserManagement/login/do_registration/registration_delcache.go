package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

// 清除缓存
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
	link, err2 := mylink.NewredisLink(0)
	if err2 != nil {
		outcometmp.Output = logs + " redis Link err"
		return err2, outcometmp
	}
	outcometmp.Output = logs + "ok"
	outcometmp.Bitmap = 1
	outcometmp.Goon = true
	outcometmp.Nextinput = string2
	link.Client.HDel(link.Ctx, "email_username", email)
	return nil, outcometmp
}
