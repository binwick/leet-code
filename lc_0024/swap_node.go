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

func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{}
	out := pre
	pre.Next = head
	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := a.Next
		pre.Next, b.Next, a.Next = b, a, b.Next
		//a.Next = b.Next
		//b.Next = a
		//pre.Next = b
		pre = a
	}

	return out.Next
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
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	traversing(head)
	out := swapPairs(head)
	traversing(out)
}
