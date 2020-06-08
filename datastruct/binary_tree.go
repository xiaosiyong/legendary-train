package datastruct

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//深度优先--前序遍历
func (T *TreeNode) PreOrderPrint() {
	t := T
	if t != nil {
		fmt.Println(t.Val)
		l := t.Left
		l.PreOrderPrint()
		r := t.Right
		r.PreOrderPrint()
	}
}

//深度优先--后序遍历
func (T *TreeNode) PostOrderPrint() {
	t := T
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
func (T *TreeNode) InOrderPrint() {
	t := T
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
func (T *TreeNode) BFSPrint() [][]int {
	var result [][]int
	if T != nil {
		dfsHelper(0, T, &result)
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

func (T *TreeNode) Insert(data int) {
	t := T
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

func (T *TreeNode) Find(data int) *TreeNode {
	t := T
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

func (T *TreeNode) Delete(data int) {
	p := T
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

	} else if pp.Left == p {
		pp.Left = child
	} else {
		pp.Right = child
	}
}
