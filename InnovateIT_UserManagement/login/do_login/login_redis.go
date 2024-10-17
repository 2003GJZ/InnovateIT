package do_login

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

// phone $ password
// xxxx$hjjhjjh$
func Login_redis_phone(string2 string) (error, string, string, byte, bool) { //查redis

	phone, s, err2 := tool.SplitString(string2, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}
	password, s2, err2 := tool.SplitString(s, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}

	link, err2 := mylink.NewredisLink(0)
	if err2 != nil {
		return err2, "", "", 0, false
	}
	var passwordMd5 string
	// 使用HGet获取哈希表的字段值
	link.Client.HGet(link.Ctx, "login_phone", phone).Scan(&passwordMd5)
	if passwordMd5 == "" {

		return nil, "", phone + "$" + password + "$" + s2, 0, true //缓存无找数据库

	}

	compareMD5 := tool.CompareMD5(password, passwordMd5)
	if compareMD5 {
		return nil, "ok", "", 1, false
	} else {
		return nil, "", "", 0, false
	}

}
