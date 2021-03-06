package algorithm

import "fmt"

const total = 13

///约瑟夫环的问题 total个数，从from开始，到count之后，count出，
///直至所有数都出
///1、循环链表实现
func JOSECircleByLink(total, from, count int) {
	fmt.Printf("总共%d个数，从%d开始，报到%d时重新计数，顺序：\n", total, from, count)
	type node struct {
		data int
		next *node
	}
	var head *node
	cur := &node{
		data: 0,
		next: nil,
	}
	for i := 1; i <= total; i++ {
		n := &node{
			data: i,
			next: nil,
		}
		if head == nil {
			head = n
		} else {
			cur.next = n
		}
		cur = n
	}
	cur.next = head
	cur = head
	for i := 1; i < from; i++ {
		cur = cur.next
	}
loop:
	if cur.next == cur {
		fmt.Println(cur.data)
		return
	}
	for i := 1; i < count-1; i++ {
		cur = cur.next
	}
	fmt.Println(cur.next.data)
	cur.next = cur.next.next
	cur = cur.next
	goto loop
}

/**
约瑟夫 数组实现
1、初始化数组，默认值0
2、从from开始，遇到未做过标识的数，+1报数
3、遇到count时，取出此时的数，重新开始报数
*/
func JOSECircleByArray(from, count int) {
	array := [total]int{}
	var outCount, index int
	for outCount < total {
		for i := from; i <= total; i++ {
			if array[i-1] == 0 {
				index++
				if index == count {
					outCount++
					array[i-1] = 1
					fmt.Println("PoP Number:", i)
					index = 0
				}
			}

		}
		from = 1
	}
}
