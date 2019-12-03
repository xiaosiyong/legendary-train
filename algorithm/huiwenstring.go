package algorithm

/*
1、单链表存储字符串，判断是否是回文字符串
2、通过快慢指针确定中间位置
3、比较前后是否相等
*/

type SinglyLinkedList struct {
	Data string
	Next *SinglyLinkedList
}

type Element interface{}

//将字符串转成链表
func InsertStringToList(input string) *SinglyLinkedList {
	start := &SinglyLinkedList{}
	//如果字符串长度为1 只有首节点
	if len(input) == 1 {
		return &SinglyLinkedList{
			Data: input,
			Next: nil,
		}
	}
	if len(input) > 1 {
		start.Data = input[0:1]
		p := start
		for i := 1; i < len(input); i++ {
			ptr := &SinglyLinkedList{}
			ptr.Data = input[i : i+1]
			p.Next = ptr
			p = ptr

		}
	}
	return start
}

func CheckIsHuiWenString(head *SinglyLinkedList) bool {
	if head == nil {
		return false
	}
	m, l := FindMiddle(head)
	var t = &SinglyLinkedList{}
	if l%2 == 0 { //偶数
		t = ReverseNode(m)
	} else {
		t = ReverseNode(m.Next)
	}
	return isNodeEqual(head, t)
}

//反转链表
func ReverseNode(node *SinglyLinkedList) *SinglyLinkedList {
	if node == nil || node.Next == nil {
		return node
	}
	var temp *SinglyLinkedList
loop:
	if node != nil {
		cur := node.Next
		node.Next = temp
		temp = node
		node = cur
		goto loop
	}
	return temp
}

//原链表、折半之后的链表
func isNodeEqual(node1 *SinglyLinkedList, node2 *SinglyLinkedList) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if (node2 != nil && node1 == nil) || (node1 != nil && node2 == nil) {
		return false
	}
loop:
	if node1 != nil && node2 != nil && node1.Data == node2.Data {
		node1 = node1.Next
		node2 = node2.Next
		goto loop
	}
	return node2 == nil
}

//查找链表中间节点
var FindMiddle = func(list *SinglyLinkedList) (*SinglyLinkedList, int) {
	if list == nil {
		return nil, 0
	}
	p := list
	q := list
	l := list
	i := 0
	for ; l != nil; l = l.Next {
		i++
	}

loop:
	if q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
		goto loop
	}

	return p, i
}

//将单链表转成字符串
func SingleLinkedListToString(list *SinglyLinkedList) string {
	//链表为空
	if len(list.Data) < 1 {
		return ""
	}
	//只有一个元素
	if list.Next == nil {
		return list.Data
	}
	//循环追加
	p := list
	str := p.Data
forLabel:
	if p.Next != nil {
		str = str + p.Next.Data
		p = p.Next
		goto forLabel
	}
	return str
}
