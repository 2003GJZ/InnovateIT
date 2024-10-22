package do_login

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
)

func Login_updatacache_phone(string2 string) (error, tool.Outcome) {
	logs := "Login_updatacache_phone:"
	outcometmp := tool.Outcome{
		logs, "", 0, false,
	}
	//字符串切割
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
	link, _ := mylink.GetredisLink()
	link.Client.HDel(link.Ctx, "login_phone", phone)           //删除
	link.Client.HSet(link.Ctx, "login_phone", phone, password) //更新
	outcometmp.Bitmap = 1
	outcometmp.Nextinput = s2
	outcometmp.Goon = true
	outcometmp.Output = logs + "ok"

	return nil, outcometmp //表示正确

}
