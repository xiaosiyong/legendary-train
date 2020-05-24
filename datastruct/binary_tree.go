package datastruct

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//深度优先--前序遍历
func (t *TreeNode) PreOrderPrint() {
	if t != nil {
		fmt.Println(t.Val)
		l := t.Left
		l.PreOrderPrint()
		r := t.Right
		r.PreOrderPrint()
	}
}

//深度优先--后序遍历
func (t *TreeNode) PostOrderPrint() {
	if t != nil {
		if t.Left != nil {
			l := t.Left
			l.PostOrderPrint()
		}
		if t.Right != nil {
			r := t.Right
			r.PostOrderPrint()
		}
		fmt.Println(t.Val)

	}
}

//深度优先--中序遍历
func (t *TreeNode) InOrderPrint() {
	if t != nil {
		if t.Left != nil {
			l := t.Left
			l.InOrderPrint()
		}
		fmt.Println(t.Val)
		if t.Right != nil {
			r := t.Right
			r.InOrderPrint()
		}

	}
}

func (t *TreeNode) BFSPrint() [][]int {
	var result [][]int
	if t != nil {
		dfsHelper(0, t, &result)
	}
	return result
}

func dfsHelper(l int, t *TreeNode, r *[][]int) {
	if t == nil {
		return
	}
	if len(*r) < l+1 {
		*r = append(*r, make([]int, 0))
	}
	(*r)[l] = append((*r)[l], t.Val)
	dfsHelper(l+1, t.Left, r)
	dfsHelper(l+1, t.Right, r)
}
