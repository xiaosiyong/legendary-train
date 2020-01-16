package datastruct

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) PreOrderPrint() {
	if t != nil {
		fmt.Println(t.Val)
		l := t.Left
		l.PreOrderPrint()
		r := t.Right
		r.PreOrderPrint()
	}
}

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

func (t *TreeNode) ByLevelPrint() {
	for t != nil {
		fmt.Println(t.Val)
	}
	dfsPrinter(t)
}

func dfsPrinter(t *TreeNode) {
	if t != nil {
		fmt.Println(t.Val)

	}
	if t.Left != nil {
		dfsPrinter(t.Left)
	}
	if t.Right != nil {
		dfsPrinter(t.Right)
	}

}
