package main

import fmt "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := l1
	node2 := l2
	remainder := false
	result := &ListNode{}
	node := result

	for node1 != nil || node2 != nil {
		val := 0
		node.Next = &ListNode{}
		node = node.Next

		if node1 != nil {
			val += node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			val += node2.Val
			node2 = node2.Next
		}
		if remainder {
			val++
		}

		if val >= 10 {
			remainder = true
			val -= 10
		} else {
			remainder = false
		}

		node.Val = val
	}

	if remainder {
		node.Next = &ListNode{}
		node.Next.Val++
	}

	return result.Next
}

func main() {
	var l1 ListNode
	var l2 ListNode
	var l1bis ListNode
	l1bis.Val = 2
	l1bis.Next = nil
	var l2bis ListNode
	l2bis.Val = 3
	l2bis.Next = nil
	l1.Val = 5
	l1.Next = &l1bis
	l2.Val = 5
	l2.Next = &l2bis

	node := addTwoNumbers(&l1, &l2)
	for node != nil {
		fmt.Printf("%v", node.Val)
		node = node.Next
	}
}
