package Easy

// Q21: https://leetcode.com/problems/merge-two-sorted-lists

type ListNode struct {
     Val int
     Next *ListNode
}

type LinkedList struct {
	head *ListNode
	length int
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	resultNode := &ListNode{}
	prev := resultNode


	for ; l1 != nil || l2 != nil ; {

		if l1 == nil {
			prev.Next = l2
			break
		}

		if l2 == nil {
			prev.Next = l1
			break
		}

		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}

		prev = prev.Next
	}

	return resultNode.Next

}