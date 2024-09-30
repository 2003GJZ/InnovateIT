package main

import (
	"github.com/go-redis/redis/v8"
	"log"
	"login/mylink"
	"login/tool"
)

// phone $ password
// xxxx$hjjhjjh$
func login_redis(string2 string) (error, string, string, byte, bool) { //查redis

	phone, s, err2 := tool.SplitString(string2, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}
	password, s2, err2 := tool.SplitString(s, "$")
	if err2 != nil {
		return err2, "", "", 0, false
	}

	err2, link := mylink.NewredisLink(0)
	if err2 != nil {
		return err2, "", "", 0, false
	}

	// 使用HGet获取哈希表的字段值
	passwordMd5, err := link.Client.HGet(link.Ctx, "login_phone", phone).Result()
	if err == redis.Nil {

		return nil, "", phone + "$" + password + "$" + s2, 0, true //缓存无找数据库

	} else if err != nil {
		log.Fatalf("HGET error: %v", err)

	}

	compareMD5 := tool.CompareMD5(password, passwordMd5)
	if compareMD5 {
		return nil, "ok", "", 1, false
	} else {
		return nil, "", "", 0, false
	}

}
