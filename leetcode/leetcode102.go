package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintByLevel(t *TreeNode) [][]int {
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
