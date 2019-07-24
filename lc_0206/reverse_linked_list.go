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

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

func traversing(node *ListNode) {
	var p = node
	for p != nil {
		fmt.Printf("%v -> ", p.Val)
		p = p.Next
	}
	fmt.Println("nil")
}

func main() {
	// 1 -> 2 -> 3 -> 4 -> 5 -> null
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	traversing(head)
	out := reverseList(head)
	traversing(out)
}
