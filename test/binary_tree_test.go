package test

import (
	"fmt"
	"legendaryTrain/datastruct"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	t4 := &datastruct.TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	t5 := &datastruct.TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	t6 := &datastruct.TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	t7 := &datastruct.TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	t2 := &datastruct.TreeNode{
		Val:   2,
		Left:  t6,
		Right: t7,
	}
	t3 := &datastruct.TreeNode{
		Val:   3,
		Left:  t4,
		Right: t5,
	}
	t1 := &datastruct.TreeNode{
		Val:   1,
		Left:  t2,
		Right: t3,
	}
	t1.PreOrderPrint()
	fmt.Println("-----------------")
	t1.PostOrderPrint()
	fmt.Println("-----------------")
	t1.InOrderPrint()
	fmt.Println("-----------------")
	t1.PreOrderPrint()
}
