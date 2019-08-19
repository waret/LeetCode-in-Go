package main

import "fmt"

// ListNode 是链接节点
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

func ss2ls(numss [][]int) []*ListNode {
	res := []*ListNode{}

	for _, nums := range numss {
		res = append(res, s2l(nums))
	}

	return res
}

func merge(lists []*ListNode) *ListNode {
	length := len(lists)
	half := length / 2

	if length == 1 {
		return lists[0]
	}

	if length == 2 {
		var (
			l0, l1   = lists[0], lists[1]
			res, cur *ListNode
		)

		if l0 == nil {
			return l1
		}
		if l1 == nil {
			return l0
		}

		if l0.Val < l1.Val {
			res, cur, l0 = l0, l0, l0.Next
		} else {
			res, cur, l1 = l1, l1, l1.Next
		}

		for l0 != nil && l1 != nil {
			if l0.Val < l1.Val {
				cur.Next, l0 = l0, l0.Next
			} else {
				cur.Next, l1 = l1, l1.Next
			}
			cur = cur.Next
		}

		if l0 != nil {
			cur.Next = l0
		}
		if l1 != nil {
			cur.Next = l1
		}

		return res
	}

	return mergeKLists([]*ListNode{mergeKLists(lists[half:]), mergeKLists(lists[:half])})
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return merge(lists)
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	head := ListNode{}
	head.Next = lists[0]
	for _, r := range lists[1:] {
		if r == nil {
			continue
		}
		l := head.Next
		cur := &head
		for l != nil && r != nil {
			if l.Val <= r.Val {
				cur.Next = l
				l = l.Next
			} else {
				cur.Next = r
				r = r.Next
			}
			cur = cur.Next
		}
		if l != nil {
			cur.Next = l
		}
		if r != nil {
			cur.Next = r
		}
	}
	return head.Next
}

func mergeKLists3(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		l0, l1 := lists[0], lists[1]
		var res, cur *ListNode
		if l0 == nil {
			return l1
		}
		if l1 == nil {
			return l0
		}
		if l0.Val < l1.Val {
			res, cur, l0 = l0, l0, l0.Next
		} else {
			res, cur, l1 = l1, l1, l1.Next
		}
		for l0 != nil && l1 != nil {
			if l0.Val < l1.Val {
				cur.Next = l0
				l0 = l0.Next
			} else {
				cur.Next = l1
				l1 = l1.Next
			}
			cur = cur.Next
		}
		if l1 != nil {
			cur.Next = l1
		}
		if l0 != nil {
			cur.Next = l0
		}
		return res
	}
	half := len(lists) / 2
	return mergeKLists3([]*ListNode{mergeKLists3(lists[0:half]), mergeKLists3(lists[half:])})
}

func main() {
	fmt.Println(l2s(mergeKLists3(ss2ls([][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}))))
	fmt.Println(l2s(mergeKLists3(ss2ls([][]int{
		{},
		{1},
	}))))
}
