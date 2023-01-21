package reverse_linked_list

func reverseLinkList_stack(head *listNode) *listNode {
	if head == nil {
		return nil
	}
	var st []int
	for {
		if head == nil {
			break
		}
		st = append(st, head.Val)
		head = head.Next
	}
	result := &listNode{}
	dummy := &listNode{}
	result.Next = dummy
	for len(st) > 0 {
		val := st[len(st)-1]
		st = st[0 : len(st)-1]
		dummy.Val = val
		if len(st) > 0 {
			dummy.Next = &listNode{}
		}
		dummy = dummy.Next
	}
	return result.Next
}
