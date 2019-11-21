package DataStruct

import (
	"github.com/cheekybits/genny/generic"
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
	next *Node
}

type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}


func (list *ItemLinkedList) Append(t Item) {
	list.lock.Lock()
	node := &Node{content:t}
	if list.head == nil {
		list.head = node
	}else{
		cur := list.head
		if cur.next != nil {
			cur = cur.next
		}
		cur.next = node
	}
	list.size++
	list.lock.Unlock()
}

func (list *ItemLinkedList) Insert(i int,item Item) error {
	list.lock.Lock()
	defer list.lock.Lock()
	node := &Node{content:item}
	if i == 1 {
		node.next = list.head
	}
	return nil
}

func (list *ItemLinkedList) Reverse(){
	p := list.head
	q := list.head.next
	list.head.next = nil
	if q != nil {
		r := q.next
		q.next = p
		p = q
		q = r
	}
	list.head = p
}


