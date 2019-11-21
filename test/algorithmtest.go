package main

import (
	"fmt"
	"legendaryTrain/AlgorithmLearn"
)

func main() {
	o1 := "1234567654321"
	o2 := "abcd11dcba"
	node1 := AlgorithmLearn.InsertStringToList(o1)
	node2 := AlgorithmLearn.InsertStringToList(o2)
	//校验是否是回文字符串
	fmt.Println(o1, "是回文字符串吗？结果：", AlgorithmLearn.CheckIsHuiWenString(node1))
	fmt.Println(o2, "是回文字符串吗？结果：", AlgorithmLearn.CheckIsHuiWenString(node2))
	//校验链表转string
	nodeString := AlgorithmLearn.SingleLinkedListToString(node1)
	fmt.Println(nodeString)
	//获取中间节点
	m, l := AlgorithmLearn.FindMiddle(node1)
	mString := AlgorithmLearn.SingleLinkedListToString(m)
	fmt.Println(mString, l)
	//中间节点反转
	r := AlgorithmLearn.ReverseNode(m)
	rString := AlgorithmLearn.SingleLinkedListToString(r)
	fmt.Println(rString)

	fmt.Println("开始测试队列：")
	q := AlgorithmLearn.InitQueue(9)
	q.Enqueue("001")
	q.Enqueue("002")
	q.Enqueue("003")
	q.Enqueue("004")
	q.Enqueue("005")
	q.Enqueue("006")
	q.Enqueue("007")

loop:
	if s := q.Dequeue(); s != "" {
		fmt.Println("Dequeue Item:", s)
		goto loop
	}

	fmt.Println("队列测试结束")
	fmt.Println("总共10个台阶，共有走法：", AlgorithmLearn.CalculateTotal1(10))
	fmt.Println("总共10个台阶，共有走法：", AlgorithmLearn.CalculateTotal2(10))
	fmt.Println("总共10个台阶，共有走法：", AlgorithmLearn.CalculateTotal3(10))

}
