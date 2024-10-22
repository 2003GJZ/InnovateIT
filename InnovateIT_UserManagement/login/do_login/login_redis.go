package do_login

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

// phone $ password
// xxxx$hjjhjjh$
func Login_redis_phone(string2 string) (error, tool.Outcome) { //查redis
	logs := "Login_redis_phone:"
	outcometmp := tool.Outcome{
		logs, "", 0, false,
	}
	phone, s, err2 := tool.SplitString(string2, "$")
	if err2 != nil {
		outcometmp.Output = logs + "SplitStringERR"
		return err2, outcometmp
	}
	password, s2, err2 := tool.SplitString(s, "$")
	if err2 != nil {
		outcometmp.Output = logs + "SplitStringERR"
		return err2, outcometmp
	}

	link, err2 := mylink.NewredisLink(0)
	if err2 != nil {
		outcometmp.Output = logs + "SplitStringERR"
		return err2, outcometmp
	}
	var passwordMd5 string
	// 使用HGet获取哈希表的字段值
	link.Client.HGet(link.Ctx, "login_phone", phone).Scan(&passwordMd5)
	if passwordMd5 == "" {
		outcometmp.Goon = true
		outcometmp.Nextinput = phone + "$" + password + "$" + s2
		outcometmp.Output = logs + "Redis is "

		return nil, outcometmp //缓存无找数据库

	}

	compareMD5 := tool.CompareMD5(password, passwordMd5)
	if compareMD5 {
		outcometmp.Output = logs + "Redis is ok"
		outcometmp.Bitmap = 1
		return nil, outcometmp
	} else {
		outcometmp.Output = logs + "password is not  right"
		return nil, outcometmp
	}

}
