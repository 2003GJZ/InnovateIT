package main

import (
	"fmt"
	"log"
)

type Liability_Node struct { //责任链节点
	dosth func(string) (error, string, string, byte, bool) //处理函数  返回值：错误，结果，传给下一个节点参数，责任节点验证失败置为0,是否继续执行,true继续，false终止

	//短链接跳转指针
	next *Liability_Node //下一个节点
	last *Liability_Node //上一个节点
}

type Liabilitylist struct {
	initial         int    //初始长度
	count           int    //责任链节点个数
	bytes           []byte //责任判断链，最大限制链路长度为1024
	liability_Nodes []Liability_Node
}

func NewLiabilitylist(initial int) *Liabilitylist {
	//返回开辟的内存
	return &Liabilitylist{
		bytes:           make([]byte, initial),           //责任节点验证失败置为0
		liability_Nodes: make([]Liability_Node, initial), //初始化责任链,初始长度
		//切片动态扩容
	}

}

func NewLiabilitylist_count(count int) *Liabilitylist {
	//返回开辟的内存
	return &Liabilitylist{
		bytes:           make([]byte, count),           //责任节点验证失败置为1
		liability_Nodes: make([]Liability_Node, count), //初始化责任链,初始长度
	}
}

func (root *Liabilitylist) AddNode(dosth func(string) (error, string, string, byte, bool)) {
	if dosth != nil {
		if root.count < root.initial {

			root.liability_Nodes[root.count] = Liability_Node{dosth: dosth}

		} else {

			root.liability_Nodes = append(root.liability_Nodes, Liability_Node{dosth: dosth})

		}

		root.count++
	} else {
		log.Println("dosth is nil")
	}
}

func (root *Liabilitylist) RunNodeList(ags string, result_partition string) (error, string, []byte) {
	var outcomes string //结果集合，临时存储待处理字符串
	var err error
	var b byte
	var goon bool
	tmp := ags //处理完后，把内容更新传递给下一个   责任节点1任务$责任节点2任务$责任节点3任务$...

	//运行节点
	for i := 0; i < root.count; i++ { // 只遍历有效节点
		var outcome string
		if tmp == "" {
			return fmt.Errorf("string exception"), outcomes, root.bytes[:len(root.bytes)] // 返回已使用的部分
		}

		if root.liability_Nodes[i].dosth != nil {
			err, outcome, tmp, b, goon = root.liability_Nodes[i].dosth(tmp)
			if err != nil {
				return err, outcomes, nil
			} else {
				outcome = outcome + result_partition
				outcomes += outcome
				root.bytes = append(root.bytes, b) // 动态添加字节
			}
			if !goon {
				break
			}

		}
	}

	return nil, outcomes, root.bytes
}
