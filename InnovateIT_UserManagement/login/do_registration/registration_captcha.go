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

	// 邮件服务器配置
	mail := gomail.NewMessage()
	mail.SetHeader("From", from)

	mail.SetHeader("To", to)
	mail.SetHeader("Subject", "验证码")
	mail.SetBody("text/html", fmt.Sprintf("你的验证码是：<b>%s</b>", code))

	d := gomail.NewDialer("smtp.qq.com", 587, from, password)

	// 发送邮件
	if err := d.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}

// 验证码发送
func Captcha_email(ags string) (error, string, string, byte, bool) {

	log := "Captcha_email:"
	//先查redis看是否存在
	link, _ := mylink.NewredisLink(0)
	email, _, _ := tool.SplitString(ags, "$")
	Captcha := link.Client.HGet(link.Ctx, "email_captcha", email)
	if Captcha.Err() == nil {
		//表示存在，则验证验证码
		log += "验证码存在"
		return nil, "", "", 0, false
	}

}

//
