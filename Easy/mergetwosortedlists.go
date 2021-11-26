package Easy


type ListNode struct {
     Val int
     Next *ListNode
}

type LinkedList struct {
	head *ListNode
	length int
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	n := ListNode{}

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		n.Next = list1
	}
	n.Next = list2








	return &n
}