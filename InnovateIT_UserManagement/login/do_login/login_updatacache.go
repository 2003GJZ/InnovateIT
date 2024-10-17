package do_login

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

func Login_updatacache_phone(string2 string) (error, string, string, byte, bool) {
	//字符串切割
	phone, s, err2 := tool.SplitString(string2, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}
	password, s2, err2 := tool.SplitString(s, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}
	link, _ := mylink.GetredisLink()
	link.Client.HDel(link.Ctx, "login_phone", phone)           //删除
	link.Client.HSet(link.Ctx, "login_phone", phone, password) //更新
	return nil, "ok", s2, 1, true                              //表示正确

}
