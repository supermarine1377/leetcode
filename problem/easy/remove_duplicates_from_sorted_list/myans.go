package remove_duplicates_from_sorted_list

import "github.com/supermarine1377/leetcode/types"

type ListNode = types.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev := head
	tail := head
	for tail != nil {
		if prev != tail && prev.Val == tail.Val {
			prev.Next = tail.Next
		} else {
			prev = tail
		}
		tail = tail.Next
	}
	return head
}
