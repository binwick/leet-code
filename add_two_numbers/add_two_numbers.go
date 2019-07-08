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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var l1p = l1
	var l2p = l2
	var l3 = &ListNode{}
	var l3p = l3
	for {
		if l1p == nil && l2p == nil {
			break
		}
		if l1p == nil {
			l1p = &ListNode{}
		}
		if l2p == nil {
			l2p = &ListNode{}
		}

		l3p.Val += l1p.Val + l2p.Val
		if l1p.Next != nil || l2p.Next != nil {
			l3p.Next = &ListNode{}
		}
		if l3p.Val >= 10 {
			l3p.Next = &ListNode{}
			l3p.Val = l3p.Val - 10
			l3p.Next.Val = 1
		}
		//
		l1p = l1p.Next
		l2p = l2p.Next
		l3p = l3p.Next
	}
	return l3
}

func traversing(node *ListNode) {
	var p = node
	for p != nil {
		fmt.Printf("%v\t", p.Val)
		p = p.Next
	}
	fmt.Println()
}

func main() {

	//l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	//l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	//l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}

	l1 := &ListNode{Val: 5}
	l2 := &ListNode{Val: 5}

	l3 := addTwoNumbers(l1, l2)
	traversing(l1)
	traversing(l2)
	traversing(l3)
}
