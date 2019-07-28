package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	node4 := &ListNode{Val: 5}
	node3 := &ListNode{Val: 4, Next: node4}
	node2 := &ListNode{Val: 3, Next: node3}
	node1 := &ListNode{Val: 2, Next: node2}
	head := &ListNode{Val: 1, Next: node1}
	// 制造一个环
	node4.Next = node2

	out := hasCycle(head)
	fmt.Println(out)
}
