package main

import (
	"fmt"
	"log"
	"login/mylink"
	"login/tool"
)

func init() {
	mylink.FileConfigJSON("/home/gjz/项目/go/login/configure.json") //初始化配置文件
}

func main() {

	err2, link := mylink.NewredisLink(0)
	if err2 != nil {
		log.Fatalf("HGET error: %v", err2)
	}

	link.Client.HSet(link.Ctx, "login_phone", "123456789", tool.GetMd5("123456789"))

	root := NewLiabilitylist(0)
	root.AddNode(login_redis)
	err2, s, bytes := root.RunNodeList("123456789$123456789$123456789", "%")
	if err2 != nil {
		log.Fatalf("HGET error: %v", err2)
	}

	//打印
	fmt.Println(s)
	fmt.Printf("%v\n", bytes)

}

//func login_mysql(string2 string) (error, string, string, byte, bool) { //查询数据库
//	phone, s, err2 := SplitString(string2, "$")
//	if err2 != nil {
//		return err2, "", "", 1, false
//	}
//	password, s2, err2 := SplitString(s, "$")
//	if err2 != nil {
//		return err2, "", "", 1, false
//	}
//
//	//TODO 查询数据库
//}
//
//func login_2(string2 string) (error, string, string, byte, bool) { //更新缓存
//	//TODO 更新缓存
//}

//func xxxx(age string) (error, string, string, byte) {
//	splitString, remaining, err := SplitString(age, ";;;")
//	if err != nil {
//		return err, "", "", 1
//	} else {
//		handle := splitString + "ok"
//		fmt.Println("处理：", handle)
//		fmt.Println("剩余: ", remaining)
//		return nil, handle, remaining, 1
//	}
//}
//
//func xxx0(age string) (error, string, string, byte) {
//	splitString, remaining, err := SplitString(age, ";;;")
//	if err != nil {
//		return err, "", "", 1
//	} else {
//		handle := splitString + string(len(splitString)/2)
//		fmt.Println("处理：", handle)
//		fmt.Println("剩余: ", remaining)
//		return nil, handle, remaining, 0
//	}
//}
//
//func xxx1(age string) (error, string, string, byte) {
//	splitString := age
//	var xx string
//
//	handle := "处理完成" + string(len(splitString)/1)
//	fmt.Println("处理：", handle)
//	return nil, handle, xx, 0
//
//}
