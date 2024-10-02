package main

import (
	"fmt"
	"log"
	"login/mylink"
	"login/tool"
	"testing"
)

func TestNewLiabilitylist(t *testing.T) {
	link, err2 := mylink.NewredisLink(0)
	if err2 != nil {
		log.Fatalf("HGET error: %v", err2)
	}

	link.Client.HSet(link.Ctx, "login_phone", "123456789", tool.GetMd5("123456789"))

	root := NewLiabilitylist(8)
	root.AddNode(login_redis)

	err2, s, bytes := root.RunNodeList("123456789$123456789$123456789", "%")
	if err2 != nil {
		log.Fatalf("HGET error: %v", err2)
	}

	//打印
	fmt.Println(s)
	fmt.Printf("%v\n", bytes)

}
