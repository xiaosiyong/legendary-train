package test

import (
	"fmt"
	"legendaryTrain/datastruct"
	"testing"
)

func TestAppend(t *testing.T) {
	link := datastruct.InitItemLinkedList()
	link.Append("1")
	link.Append("2")
	link.Append("3")
	link.Append("4")
	link.Append("5")
	//fmt.Println(datastruct.IsPalindrome(link))
	link.Append("6")
	link.Append("6")
	link.Append("5")
	link.Append("4")
	link.Append("3")
	link.Append("2")
	link.Append("1")
	fmt.Println(datastruct.IsPalindrome(link))

	//link.Print()
	//link.Append("6")
	//link.Print()
	//fmt.Println(link.IndexOf(2))
	//link.Insert(3, "8")
	//link.Print()
	//link.Reverse()
	//link.Print()
	//link.RemoveAt(3)
	//link.Print()
	//
	//link1 := datastruct.InitItemLinkedList()
	//link1.Append("1")
	//link1.Append("3")
	//link1.Append("5")
	//link1.Append("7")
	//link2 := datastruct.InitItemLinkedList()
	//link2.Append("2")
	//link2.Append("4")
	//link2.Append("10")
	//link2.Append("16")
	//link2.Append("17")
	//link2.Append("18")
	//link2.Append("28")
	//
	//m := datastruct.MergeSortedLinkList(link1, link2)
	//m.Print()
	//m.DeleteReverseNNode(2)
	//m.Print()
	//middle := m.FindMiddleNode()
	//fmt.Println(middle)
}
