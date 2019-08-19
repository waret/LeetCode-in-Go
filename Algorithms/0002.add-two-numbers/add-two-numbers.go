package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers 数字相加
// 1. 循环遍历两个加数串和进位值，进入下一位运算
// 2. 第一个Node节点的处理
// 3. 进位的持久化
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	carray := 0

	for l1 != nil || l2 != nil || carray > 0 {
		if l1 != nil {
			carray += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carray += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: carray % 10}
		carray = carray / 10
		cur = cur.Next
	}
	return result.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) (result *ListNode) {
	//var result *ListNode
	var cur *ListNode
	carray := 0

	for l1 != nil || l2 != nil || carray > 0 {
		if cur == nil {
			// 个位
			cur = &ListNode{}
		} else {
			// 后续位
			cur.Next = &ListNode{}
			cur = cur.Next
		}
		if result == nil {
			result = cur
		}
		if l1 != nil {
			carray += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carray += l2.Val
			l2 = l2.Next
		}
		cur.Val = carray % 10
		carray = carray / 10
	}
	return
}
