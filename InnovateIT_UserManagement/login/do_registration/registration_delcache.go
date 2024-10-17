package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

// 清除缓存
func Delcache_email(string2 string) (error, string, string, byte, bool) {
	email, _, err2 := tool.SplitString(string2, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}
	link, err2 := mylink.NewredisLink(0)
	if err2 != nil {
		return err2, "", "", 0, false
	}
	log := "Delcache_email:ok"
	link.Client.HDel(link.Ctx, "email_username", email)
	return nil, log, string2, 1, true
}
