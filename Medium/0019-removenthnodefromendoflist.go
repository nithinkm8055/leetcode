package Medium

type Pair struct {
	curr *ListNode
	prev *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	positionalMap := make(map[int]Pair)

	temp := head
	counter := 1

	var prev *ListNode

	for temp != nil {
		counter++
		positionalMap[counter] = Pair{
			temp,
			prev,
		}
		prev = temp
		temp = temp.Next
	}

	positionalNodeToRemove := n - 1

	pair := positionalMap[counter-positionalNodeToRemove]
	if pair.prev != nil {
		pair.prev.Next = pair.curr.Next
	} else {
		head = pair.curr.Next
	}

	return head
}
