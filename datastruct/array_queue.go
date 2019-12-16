package datastruct

type ArrayQueue struct {
	Size  int
	Items []string
	Head  int
	Tail  int
}

type CircleQueue struct {
	Size  int
	Items []string
	Head  int
	Tail  int
}

func InitQueue(l int) *ArrayQueue {
	return &ArrayQueue{
		Size:  l,
		Items: make([]string, l),
	}
}

//入队
func (a *ArrayQueue) Enqueue(s string) bool {
	if a.Tail == a.Size {
		if a.Head == 0 {
			return false
		}
		for i := a.Head; i < a.Tail; i++ {
			a.Items[i-a.Head] = a.Items[i]
		}
		a.Tail -= a.Head
		a.Head = 0
	}
	a.Items[a.Tail] = s
	a.Tail++
	return true
}

//出队
func (a *ArrayQueue) Dequeue() string {
	if a.Head == a.Tail {
		return ""
	}
	s := a.Items[a.Head]
	a.Head++
	return s
}

func InitCircleQueue(l int) *CircleQueue {
	return &CircleQueue{
		Size:  l,
		Items: make([]string, l),
	}
}

func (a *CircleQueue) Enqueue(s string) bool {
	if (a.Tail+1)%a.Size == a.Head {
		return false
	}
	a.Items[a.Tail] = s
	a.Tail = (a.Tail + 1) % a.Size
	return true
}

func (a *CircleQueue) Dequeue() string {
	if a.Tail == a.Head {
		return ""
	}
	str := a.Items[a.Head]
	a.Head = (a.Head + 1) % a.Size
	return str
}
