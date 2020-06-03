package main

/*
面试题 02.04. 分割链表

编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。
分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。

*/
func main() {

}

/*
先创建两个头结点分别用于连接小于部分和大于等于部分
遍历该链表，若当前结点小于x,将其插到小于的链表上否则插到大于链表上。
遍历完毕后，将大于等于部分链表的连到小于部分的后面即可.

*/
func partition(head *ListNode, x int) *ListNode {
	// 两个头节点
	ltHead := &ListNode{}
	gtHead := &ListNode{}
	// 游标
	ltCursor := ltHead
	gtCursor := gtHead
	c := head

	for c != nil {
		if c.Val < x {
			// 小于则插入小于链表
			ltCursor.Next = c
			ltCursor = ltCursor.Next
		} else {
			// 大于则插入大于链表
			gtCursor.Next = c
			gtCursor = gtCursor.Next
		}
		c = c.Next
	}
	// 把大于链表链接到小于链表尾部
	ltCursor.Next = gtHead.Next
	// 同时清空大于链表的尾节点
	gtCursor.Next = nil
	return ltHead.Next
}
