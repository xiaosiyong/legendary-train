package test

import (
	"fmt"
	"legendaryTrain/algorithm"
	"testing"
)

func TestCheckIsHuiWenString(t *testing.T) {
	o1 := "1234567654321"
	o2 := "abcd11dcba"
	node1 := algorithm.InsertStringToList(o1)
	node2 := algorithm.InsertStringToList(o2)
	//校验是否是回文字符串
	fmt.Println(o1, "是回文字符串吗？结果：", algorithm.CheckIsHuiWenString(node1))
	fmt.Println(o2, "是回文字符串吗？结果：", algorithm.CheckIsHuiWenString(node2))

	//校验链表转string
	nodeString := algorithm.SingleLinkedListToString(node1)
	fmt.Println(nodeString)
	//获取中间节点
	m, l := algorithm.FindMiddle(node1)
	mString := algorithm.SingleLinkedListToString(m)
	fmt.Println(mString, l)
	//中间节点反转
	r := algorithm.ReverseNode(m)
	rString := algorithm.SingleLinkedListToString(r)
	fmt.Println(rString)
}
