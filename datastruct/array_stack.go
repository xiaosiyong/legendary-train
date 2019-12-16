package datastruct

//数组实现的栈
type ArrayStack struct {
	Length int
	Count  int
	Items  []string
}

func InitArrayStack(n int) *ArrayStack {
	return &ArrayStack{
		Length: n,
		Count:  0,
		Items:  make([]string, n),
	}
}

func (a *ArrayStack) Push(item string) bool {
	if a.Count == a.Length {
		return false
	}
	a.Items[a.Count] = item
	a.Count++
	return true
}

func (a *ArrayStack) Pop() string {
	if a.Count == 0 {
		return ""
	}
	item := a.Items[a.Count-1]
	a.Count--
	return item
}
