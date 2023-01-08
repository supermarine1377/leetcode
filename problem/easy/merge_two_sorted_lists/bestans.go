package merge_two_sorted_lists

import "github.com/supermarine1377/leetcode/types"

type ListNode = types.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy ListNode
	tail := &dummy

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			tail.Next = list2
			list2 = list2.Next
		} else {
			tail.Next = list1
			list1 = list1.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}

	return dummy.Next
}

func mergeTwoListsRecursively(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		if list1 != nil {
			return list1
		} else {
			return list2
		}
	}
	// case: list1 != nil && list2 != nil
	if list1.Val > list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	} else {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
}
