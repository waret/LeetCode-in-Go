package problem0024

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode is definition for singly-linked list
type ListNode = kit.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}


func swapPairs2(head *ListNode) *ListNode {
	cur := &ListNode{Next:head}
	old := cur
	for cur.Next != nil && cur.Next.Next != nil {
		second := cur.Next.Next
		cur.Next.Next = second.Next
		second.Next = cur.Next
		cur.Next = second
		cur = cur.Next.Next
	}
	return old.Next
}

func swapPairs3(head *ListNode) *ListNode {
	cur := &ListNode{Next:head}
	old := cur
	for cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		second := first.Next
		last := second.Next
		cur.Next = second
		second.Next = first
		first.Next = last
		cur = cur.Next.Next
	}
	return old.Next
}