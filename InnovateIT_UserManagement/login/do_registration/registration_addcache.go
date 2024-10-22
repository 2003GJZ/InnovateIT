package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

// 加添缓存，查看是否有该用户
func Addcache_email(string2 string) (error, string, string, byte, bool) {
	email, s, err2 := tool.SplitString(string2, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}
	usernameMd5, s2, err2 := tool.SplitString(s, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}
	log := "Addcache_email:"
	goon := true
	link, _ := mylink.NewredisLink(0)
	link.Client.HSet(link.Ctx, "email_username", email, usernameMd5)
	log += "ok"
	if usernameMd5 == "NULL" {
		goon = false
		s2 = email + "$" + s2
	}
	return nil, log, s2, 1, goon
}
