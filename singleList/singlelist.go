package main

import "fmt"

/*
单链表相关问题：
1）判断是否有环：快慢指针，能够重合证明存在环
2）如果有环，找出环的入口 ：
      推导得出从相遇点到入口，和从起点到入口距离相同，两个指针，相遇时即为入口
3) 求环的长度：慢指针从相遇点再走一圈
4）单链表逆序
*/


type node struct {
	val int
	next *node
}

//判断是否有环
func isIntersect(head *node) bool {
	fast := head
	slow := head
	for  {
		if fast == nil || slow == nil{
			break
		}
		fast = fast.next.next
		slow = slow.next
		if slow == fast{
			return true
		}
	}
	return false
}

//如果有环，找出环的入口
func entrance(head *node)*node{
	fast := head
	slow := head
	for  {
		if fast == nil || slow == nil{
			break
		}
		fast = fast.next.next
		slow = slow.next
		if slow == fast{
			//走到了相遇点
			enter := head
			for {
				enter = enter.next
				slow = slow.next
				if enter == slow{
					return enter
				}
			}
		}
	}
	return nil
}

//求环的长度
func ringLen(head *node)int{
	fast := head
	slow := head
	for  {
		if fast == nil || slow == nil{
			break
		}
		fast = fast.next.next
		slow = slow.next
		if slow == fast{
			//走了了相遇点
			temp := slow
			len := 0
			for  {
				slow = slow.next
				len++
				if slow == temp {
					return len
				}
			}
		}
	}
	return 0
}

func main(){
	var node1,node2,node3,node4,node5 node

	//构建一个循环单链表
	node1 = node{
		val : 1,
		next: &node2,
	}
	node2 = node{
		val : 2,
		next: &node3,
	}
	node3 = node{
		val : 3,
		next: &node4,
	}
	node4 = node{
		val : 4,
		next: &node5,
	}
	node5 = node{
		val : 5,
		next: &node2,
	}
	head := &node1
	res := isIntersect(head)
	if res{
		inter := entrance(head)
		fmt.Println(inter.val)
	}
	len := ringLen(head)
	fmt.Println(len)

}