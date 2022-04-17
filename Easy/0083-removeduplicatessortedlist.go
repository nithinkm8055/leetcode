package Easy


//type ListNode struct {
//	Val int
//	Next *ListNode
//}

func deleteDuplicates(head *ListNode) *ListNode {

	result := new(ListNode)
	prev := result
	for head != nil {

		if head.Next != nil && head.Val != head.Next.Val || head.Next == nil{
			prev.Next = head
			prev = prev.Next
		}
		head = head.Next
	}
	return result.Next

}