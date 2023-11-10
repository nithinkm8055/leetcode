package Medium

type listNode struct {
	Val  int
	Next *listNode
}

func swapPairs(head *listNode) *listNode {
	if head == nil || head.Next == nil {
		return head
	}

	firstNode := head
	secondNode := firstNode.Next

	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}
