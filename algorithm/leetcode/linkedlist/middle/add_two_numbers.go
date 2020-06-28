package main

import (
	"fmt"
)

/*
面试题 02.05. 链表求和 中等
给定两个用链表表示的整数，每个节点包含一个数位。
这些数位是反向存放的，也就是个位排在链表首部。
编写函数对这两个整数求和，并用链表形式返回结果。

输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
输出：2 -> 1 -> 9，即912
进阶：假设这些数位是正向存放的，请再做一遍。
输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
输出：9 -> 1 -> 2，即912
*/
func main() {

	l1 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Printf("%v \n", l3.Val)
		l3 = l3.Next
	}
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers
/*
动态计算 走一步加一次 有进位则当做下一轮的初始值
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	ret := &ListNode{} // 定义虚拟头节点
	cursor := ret      // 定义游标
	carry := 0         // 进位变量
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry // sum初始化为carry，首先计算进位
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		// 以sum对10取模的结果来创建新的节点
		cursor.Next = &ListNode{Val: sum % 10}
		cursor = cursor.Next // p向后移动一个节点，指向新创建的当前节点
		carry = sum / 10     // 更新进位
	}
	return ret.Next
}
