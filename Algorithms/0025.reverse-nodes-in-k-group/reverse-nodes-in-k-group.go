package main

import (
	"fmt"
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode defines for singly-linked list.
type ListNode = kit.ListNode

func main() {
	fmt.Println(kit.List2Ints(reverseKGroup3(kit.Ints2List([]int{1, 2, 3, 4, 5}), 3)))
}

func reverseKGroup3(head *ListNode, k int) *ListNode {
	if k < 2 || head == nil || head.Next == nil {
		return head
	}
	tail := head
	x := 1
	for ; x < k && tail != nil; x++ {
		tail = tail.Next
	}
	if x == k && tail != nil {
		tailNext := tail.Next
		tail.Next = nil
		curPre, cur := head, head.Next
		for cur != nil {
			curPre, cur, cur.Next = cur, cur.Next, curPre
		}
		head.Next = reverseKGroup3(tailNext, k)
		return tail
	} else {
		return head
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 || head == nil || head.Next == nil {
		return head
	}

	tail, needReverse := getTail(head, k)

	if needReverse {
		tailNext := tail.Next
		/* 斩断 tail 后的链接 */
		tail.Next = nil
		head, tail = reverse(head, tail)
		/* tail 后面接上尾部的递归处理 */
		tail.Next = reverseKGroup(tailNext, k)
	}

	return head
}

func getTail(head *ListNode, k int) (*ListNode, bool) {
	for k > 1 && head != nil {
		head = head.Next
		k--
	}
	return head, k == 1 && head != nil
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	curPre, cur := head, head.Next
	for cur != nil {
		curPre, cur, cur.Next = cur, cur.Next, curPre
	}
	return tail, head
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	cur := head
	count := 0
	for cur != nil && count < k {
		cur = cur.Next
		count++
	}
	if count == k {
		cur = reverseKGroup2(cur, k)
		for ; count > 0; count-- {
			tmp := head.Next
			head.Next = cur
			cur = head
			head = tmp
		}
		head = cur
	}
	return head
}
