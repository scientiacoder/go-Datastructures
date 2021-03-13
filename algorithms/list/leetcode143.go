package list

/*
	给定一个单链表 L：L0→L1→…→Ln-1→Ln
	将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

	你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
	示例 1:
	给定链表 1->2->3->4, 重新排列为 1->4->2->3.

	示例 2:
	给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/reorder-list
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// 找中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 反转后面
	next := slow.Next
	slow.Next = nil
	var pre *ListNode
	cur := next
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	// pre为第二个链表的首部
	// 合并链表
	ll, rr := head, pre
	for ll != nil && rr != nil {
		llNext := ll.Next
		rrNext := rr.Next
		ll.Next = rr
		rr.Next = llNext
		ll = llNext
		rr = rrNext
	}

}
