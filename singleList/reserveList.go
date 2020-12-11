package main

import (
	"fmt"
)

type Node struct {
	val int
	next *Node
}

//head 1 -> 2 ->3 ->4

func reverseList(head *Node)*Node{
	var prev,next *Node

	for head != nil{
		next = head.next
		head.next = prev
		prev = head
		head = next
	}
	return prev
}

func main(){
	var node1,node2,node3,node4,node5 Node

	//构建一个循环单链表
	node1 = Node{
		val : 1,
		next: &node2,
	}
	node2 = Node{
		val : 2,
		next: &node3,
	}
	node3 = Node{
		val : 3,
		next: &node4,
	}
	node4 = Node{
		val : 4,
		next: &node5,
	}
	node5 = Node{
		val : 5,
		next: nil,
	}
	head := &node1
	head = reverseList(head)
	ptr := head
	for  {
		if ptr == nil {
			break
		}
		fmt.Println(ptr.val,ptr.next)
		ptr = ptr.next
	}
}