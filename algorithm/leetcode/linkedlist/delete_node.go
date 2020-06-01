package linkedlist

/*
面试题 02.03. 删除中间节点
实现一种算法，删除单向链表中间的某个节点（即不是第一个或最后一个节点），假定你只能访问该节点。
输入：单向链表a->b->c->d->e->f中的节点c
结果：不返回任何数据，但该链表变为a->b->d->e->f

只能访问该节点则不能通过将prev.next指向当前节点.next来实现
则把下一节点值赋值给当前节点 然后把当前节点的下一节点换成下下个节点。
*/
func main() {

}

// 4ms 2.9M内存
func deleteNode(node *ListNode) {
	*node = *node.Next
}

// 0ms 2.9M内存
func deleteNode2(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
