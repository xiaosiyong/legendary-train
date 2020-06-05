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

//广度优先
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

func Insert(t *TreeNode, data int) {
	if t == nil {
		t = &TreeNode{Val: data}
		return
	}
	for t != nil {
		if data > t.Val {
			if t.Right == nil {
				r := &TreeNode{Val: data}
				t.Right = r
				return
			}
			t = t.Right
		} else {
			if t.Left == nil {
				l := &TreeNode{Val: data}
				t.Left = l
				return
			}
			t = t.Left
		}
	}
}

func Find(t *TreeNode, data int) *TreeNode {
	var r *TreeNode
	for t != nil {
		if data > t.Val {
			t = t.Right
		} else if data < t.Val {
			t = t.Left
		} else {
			return t
		}

	}
	return r
}

func Delete(t *TreeNode, data int) {
	p := t
	var pp *TreeNode
	for p != nil && p.Val != data {
		pp = p
		if data > p.Val {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	if p == nil { //can't find data in tree
		return
	}
	if p.Left != nil && p.Right != nil { //both left node and right node isn't nil
		minp := p.Right
		minpp := p
		for minp.Left != nil {
			minpp = minp
			minp = minp.Left
		}
		p.Val = minp.Val
		p = minp
		pp = minpp
	}

	var child *TreeNode
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	}
	if pp == nil {
		t = child
	} else if pp.Left == p {
		pp.Left = child
	} else {
		pp.Right = child
	}
}
