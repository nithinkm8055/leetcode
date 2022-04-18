package Medium

type ListNode struct {
	Val  int
	Next *ListNode
}

// Q2: https://leetcode.com/problems/add-two-numbers/

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	returnNode := &ListNode{}
	prev := returnNode
	carry := 0

	for l1 != nil || l2 != nil {

		prev.Val = l1.Val + l2.Val + carry
		l1 = l1.Next
		l2 = l2.Next

		if prev.Val > 9 {
			prev.Val = prev.Val % 10
			carry = 1
		} else {
			carry = 0
		}

		if l1 == nil && l2 != nil {
			l1 = &ListNode{}
			l1.Val = 0
		} else if l2 == nil && l1 != nil {
			l2 = &ListNode{}
			l2.Val = 0
		}

		if l1 != nil || l2 != nil {
			prev.Next = &ListNode{}
			prev = prev.Next
		}

		if l1 == nil && l2 == nil {
			if carry == 1 {
				prev.Next = &ListNode{}
				prev = prev.Next
				prev.Val = 1
			}
		}
	}
	return returnNode
}
