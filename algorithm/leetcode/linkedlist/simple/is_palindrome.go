package main

/*
面试题 02.06. 回文链表

编写一个函数，检查输入的链表是否是回文的。
输入： 1->2
输出： false

输入： 1->2->2->1
输出： true

进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
func main() {

}

/*
快慢指针 将链表的一半存到slice中
然后取出来和慢指针挨个对比
12321 slice中存放了最前面的12 慢指针指向3 判断为奇数个将慢指针后移动 指向第2个2
然后开始对比 slice中倒着取出来
如果有栈的话就可以直接取了 不过go中没有就将就slice用吧...
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast := head
	slow := head
	stack := make([]int, 0)

	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}
	// fast 不为空说明前面for循环是因为 fast.Next != nil 退出的
	// 则说明链表长度为奇数 对比之前需要把 slow 指针指向下一个
	if fast != nil {
		slow = slow.Next
	}
	l := len(stack)
	var pop int

	for slow != nil {
		l = len(stack)
		if l > 0 {
			pop = stack[l-1]
			stack = stack[:l-1]

			if pop != slow.Val {
				return false
			}
		}
		slow = slow.Next
	}
	return true
}
