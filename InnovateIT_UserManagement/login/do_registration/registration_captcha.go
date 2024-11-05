package do_registration

import (
	"InnovateIT_UserManagement/mylink"
	"InnovateIT_UserManagement/tool"
	"crypto/rand"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/big"
	"strconv"
)

// 生成6位随机验证码
func generateCode() string {
	var code string
	for i := 0; i < 6; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(10))
		code += strconv.FormatInt(num.Int64(), 10)
	}
	return code
}

// 通过QQ邮箱发送验证码，并把验证码加入，redis，ttl设置为3分钟
func sendEmail(to string, code string) error {

	// 发件人邮箱信息
	from := "1820284294@qq.com"    // QQ邮箱
	password := "dhielscygzirdfea" // 替换为你的QQ邮箱授权码
	expirationMinutes := 3         // 验证码有效期，单位为分钟

	// 邮件服务器配置
	mail := gomail.NewMessage()
	mail.SetHeader("From", from)

	mail.SetHeader("To", to)
	mail.SetHeader("Subject", "验证码")
	mail.SetBody("text/html", fmt.Sprintf(`
<html>
<head>
    <title>验证码通知</title>
</head>
<body>
    <p>尊敬的用户，您好！</p>
    <p>您正在注册“幸福年华志愿行”。</p>
    <p>您的验证码是：<strong>%s</strong></p>
    <p>请在 %d 分钟内使用此验证码完成验证。</p>
    <p>如果这不是您本人的操作，请忽略此邮件！</p>
    <p>感谢您对我们服务的使用。</p>
</body>
</html>
`, code, expirationMinutes))

	d := gomail.NewDialer("smtp.qq.com", 587, from, password)

	// 发送邮件
	if err := d.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}

// 1.4验证码发送   邮箱$
func Captcha_email_send(ags string) (error, tool.Outcome) {

	logs := "Captcha_email_send:"
	outcometmp := tool.Outcome{
		logs, "", 0, false,
	}
	//先查redis看是否存在
	link, _ := mylink.NewredisLink(0)
	email, _, err := tool.SplitString(ags, "$")
	if err != nil {
		outcometmp.Output = logs + "参数解析错误"
		outcometmp.Bitmap = 0
		outcometmp.Goon = false
		return nil, outcometmp
	}
	htable := tool.Redis_htable{
		Htabname:   "email_captcha",
		Redis_link: link,
	}
	//获取验证码
	captcha := htable.Query_caching(email)
	if captcha != "" {
		//表示存在，则无需再次发送验证码
		outcometmp.Output = logs + "验证码存在"
		outcometmp.Bitmap = 0
		outcometmp.Goon = false

	} else {
		//否则发送验证码
		//生成验证码
		code := generateCode()
		//调用验证码发送
		err := sendEmail(email, code)
		if err != nil {
			outcometmp.Output = logs + "验证码发送失败"
			outcometmp.Bitmap = 0
			outcometmp.Goon = false
		} else {
			outcometmp.Output = logs + "验证码发送成功" + " <code:" + code + ">"
			outcometmp.Bitmap = 1
			outcometmp.Goon = true
			outcometmp.Nextinput = ags
			//加入缓存
			htable.Insert_caching(email, code)
		}
	}
	return nil, outcometmp
}

// 2.1证验证码 邮箱$验证码$用户名$密码------->邮箱$用户名$密码
func Captcha_email_verify(ags string) (error, tool.Outcome) {
	logs := "Captcha_email_verify:"
	outcometmp := tool.Outcome{
		logs, "", 0, false,
	}
	//先查redis看是否存在
	link, _ := mylink.NewredisLink(0)
	email, s2, err := tool.SplitString(ags, "$")
	code, s3, err1 := tool.SplitString(s2, "$")

	if err != nil || err1 != nil {
		outcometmp.Output = logs + "参数解析错误"
		outcometmp.Bitmap = 0
		outcometmp.Goon = false
		return nil, outcometmp
	}

	htable := tool.Redis_htable{
		Htabname:   "email_captcha",
		Redis_link: link,
	}

	captcha := htable.Query_caching(email)
	if captcha == code {
		outcometmp.Output = logs + "验证码正确"
		outcometmp.Bitmap = 1
		outcometmp.Goon = true
	} else {
		outcometmp.Output = logs + "验证码错误"
		outcometmp.Bitmap = 0
		outcometmp.Goon = false
	}
	outcometmp.Nextinput = email + "$" + s3

	return nil, outcometmp
}
