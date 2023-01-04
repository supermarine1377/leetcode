package reverse_linked_list

func reverseList(head *listNode) *listNode {
	curr := head
	var prev *listNode
	for curr != nil {
		nxt := curr.next
		curr.next = prev
		prev = curr
		curr = nxt
	}
	return prev
}
