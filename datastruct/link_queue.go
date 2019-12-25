package datastruct

import "github.com/cheekybits/genny/generic"

type IntItem generic.Number

type NodeInt struct {
	data IntItem
	next *NodeInt
}

type LinkQueue struct {
	maxsize     int
	currentSize int
	head        *NodeInt
	tail        *NodeInt
}

func CreateLinkQueue(s int) *LinkQueue {
	return &LinkQueue{
		maxsize:     s,
		currentSize: 0,
		head:        nil,
		tail:        nil,
	}
}

func (l *LinkQueue) Enqueue(item IntItem) bool {
	if l.currentSize < l.maxsize {
		d := &NodeInt{
			data: item,
			next: nil,
		}
		if l.currentSize == 0 {
			l.head = d
			l.tail = d
		} else {
			l.tail.next = d
			l.tail = d
		}
		l.currentSize++
		return true
	}
	return false
}

func (l *LinkQueue) Dequeue() IntItem {
	if l.currentSize < 1 {
		return 0
	}
	v := l.head.data
	l.head = l.head.next
	l.currentSize--
	return v
}
