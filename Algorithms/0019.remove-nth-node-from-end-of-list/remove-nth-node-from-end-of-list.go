package main

import (
	"fmt"
)

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 是节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// convert *ListNode to []int
func l2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// convert []int to *ListNode
func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	d, headIsNthFromEnd := getDaddy(head, n)

	if headIsNthFromEnd {
		// 删除head节点
		return head.Next
	}

	d.Next = d.Next.Next

	return head
}

// 获取倒数第N个节点的父节点
func getDaddy(head *ListNode, n int) (daddy *ListNode, headIsNthFromEnd bool) {
	daddy = head

	for head != nil {
		if n < 0 {
			daddy = daddy.Next
		}

		n--
		head = head.Next
	}

	// n==0，说明链的长度等于n
	headIsNthFromEnd = n == 0

	return
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	daddy, last := head, head
	for last != nil {
		if n < 0 {
			daddy = daddy.Next
		}
		n--
		last = last.Next
	}
	if n == 0 {
		return head.Next
	}
	daddy.Next = daddy.Next.Next
	return head
}


func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	s := 0
	find := head
	for find != nil {
		find = find.Next
		s++
	}
	find = head
	if s == n {
		return head.Next
	}
	for s = s - n; s > 1; {
		find = find.Next
		s--
	}
	find.Next = find.Next.Next
	return head
}

func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	left, right := head, head
	for i := 0; i < n && right != nil; i++ {
		right = right.Next
	}
	if right == nil {
		return head.Next
	}
	for right.Next != nil{
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return head
}

func main() {
	fmt.Println(l2s(removeNthFromEnd1(s2l([]int{1, 2, 3, 4, 5}), 2)))
	fmt.Println(l2s(removeNthFromEnd1(s2l([]int{1, 2}), 1)))
	fmt.Println(l2s(removeNthFromEnd1(s2l([]int{1, 2}), 2)))
	fmt.Println(l2s(removeNthFromEnd1(s2l([]int{1}), 1)))
}
