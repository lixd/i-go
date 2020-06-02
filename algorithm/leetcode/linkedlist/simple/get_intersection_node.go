package main

/*
面试题 02.07. 链表相交
给定两个（单向）链表，判定它们是否相交并返回交点。请注意相交的定义基于节点的引用，而不是基于节点的值。
换句话说，如果一个链表的第k个节点与另一个链表的第j个节点是同一节点（引用完全相同），则这两个链表相交。

输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

*/
func main() {

}

/*
解析 https://www.jianshu.com/p/1a326d52a696

如果两个链表有相交的话 则让两个指针p1 p2 分别走完两个链表不同的部分和相同的部分 这样二者走的路程是一样的。
p1
↓
1->2->3↘
		 6->7->8
   4->5↗
   ↑
   p2

p1 走的路程为 123678 为 nil 后指向另一链表 head 即这里的 4 然后继续 45678
p2 则为 45678 然后切换到 1 继续 123
由于走的路程都是12345678 所以两个指针最终会相交
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB

	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
