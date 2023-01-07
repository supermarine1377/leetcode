package reverse_linked_list

import "github.com/supermarine1377/leetcode/types"

type listNode = types.ListNode

func reverseList(head *listNode) *listNode {
	curr := head
	var prev *listNode
	for curr != nil {
		nxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}
	return prev
}

func reveseListRecursively(head *listNode) *listNode {
	if head == nil {
		return nil
	}
	newHead := head
	if head.Next != nil {
		newHead = reveseListRecursively(head.Next)
		head.Next.Next = head
	}
	head.Next = nil

	return newHead
}
