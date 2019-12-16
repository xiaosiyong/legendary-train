package test

import (
	"fmt"
	"legendaryTrain/datastruct"
	"testing"
)

func TestEnqueue(t *testing.T) {
	fmt.Println("开始测试队列：")
	q := datastruct.InitQueue(9)
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
}
