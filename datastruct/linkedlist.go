package datastruct

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
	"strconv"
	"sync"
)

/*
单向线性链表
Append(t)		// 将元素 t 追加到链表尾部
Insert(i, t)    	// 在位置 i 处插入元素 t
RemoveAt(i)		// 移除位置 i 的元素
IndexOf(t)		// 返回元素 t 的位置
IsEmpty(l)		// l 是空链表则返回 true
Size()			// 返回链表的长度
String()		// 返回链表的字符串表示
Head()			// 返回链表的首结点，以便迭代链表
*/

type Item generic.Type

type Node struct {
	content Item
	next    *Node
}

type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

func InitItemLinkedList() *ItemLinkedList {
	return &ItemLinkedList{
		head: nil,
		size: 0,
		lock: sync.RWMutex{},
	}
}

func (list *ItemLinkedList) Append(t Item) {
	list.lock.Lock()
	node := &Node{content: t}
	if list.head == nil {
		list.head = node
	} else {
		cur := list.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node
	}
	list.size++
	list.lock.Unlock()
}

func (list *ItemLinkedList) Insert(i int, item Item) bool {
	list.lock.Lock()
	defer list.lock.Unlock()
	node := &Node{content: item}
	if i > list.size { //直接
		return false
	}
	if i == 1 {
		node.next = list.head
		list.head = node
		list.size++
		return true
	} else {
		cur := list.head
		for t := 1; t < i-1; t++ {
			cur = cur.next
		}
		node.next = cur.next
		cur.next = node
		list.size++
		return true
	}
}

func (list *ItemLinkedList) Reverse() {
	p := list.head
	q := list.head.next
	list.head.next = nil
	for q != nil {
		r := q.next
		q.next = p
		p = q
		q = r
	}
	list.head = p
}

func (list *ItemLinkedList) RemoveAt(i int) {
	if list.size > i {
		cur := list.head
		for t := 1; t < i-1; t++ {
			cur = cur.next
		}
		cur.next = cur.next.next
		list.size--
	}
}

func (list *ItemLinkedList) IndexOf(i int) Item {
	if list.size < i {
		return nil
	}
	t := list.head
	for j := 1; j < i; j++ {
		t = t.next
	}
	return t.content
}

func (list *ItemLinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *ItemLinkedList) Print() {
	if list.head == nil {
		fmt.Println("link is nil")
		return
	}
	if list.size == 1 {
		fmt.Println(list.head.content)
		return
	}
	temp := list.head.content.(string)
	cur := list.head
	for cur.next != nil {
		temp = temp + " " + cur.next.content.(string)
		cur = cur.next
	}
	fmt.Println("The size:[", list.size, "]", temp)
}

//TODO 检测链表中是否存在环
func CheckIsCircleInLinkList(item *ItemLinkedList) bool {
	if item.head != nil {
		quick := item.head
		slow := item.head
		for quick != nil && quick.next != nil {
			quick = quick.next.next
			slow = slow.next
			if quick == slow {
				return true
			}
		}
	}
	return false
}

//有序链表合并
func MergeSortedLinkList(item1, item2 *ItemLinkedList) *ItemLinkedList {
	result := &ItemLinkedList{}
	if item1.head != nil && item2.head != nil {
		p := item1.head
		q := item2.head
		for p != nil && q != nil {
			i1, _ := strconv.Atoi(p.content.(string))
			i2, _ := strconv.Atoi(q.content.(string))
			if i1 < i2 {
				result.Append(p.content.(string))
				p = p.next
			} else {
				result.Append(q.content.(string))
				q = q.next
			}
		}
		for p != nil {
			result.Append(p.content.(string))
			p = p.next
		}
		for q != nil {
			result.Append(q.content.(string))
			q = q.next
		}

	}
	return result
}

//删除倒数第N个节点
func (list *ItemLinkedList) DeleteReverseNNode(n int) {
	if list.size > n {
		cur := list.head
		for i := 1; i < list.size-n; i++ {
			cur = cur.next
		}
		cur.next = cur.next.next
		list.size--
	}
}

//找出中间结点
func (list *ItemLinkedList) FindMiddleNode() Item {
	if list.size < 2 {
		return list.head.content
	}
	p := list.head
	q := list.head
	for q != nil && q.next != nil {
		p = p.next
		q = q.next.next
	}
	return p.content
}
