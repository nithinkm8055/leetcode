package Medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, k int) *TreeNode {

	// search Node
	// if node does not exist return nil
	// if node is present
	// node is leaf
	// node has single child
	// node has 2 children

	parent, curr := searchNode(root, k)

	if curr == nil {
		return root
	}

	// node is leaf or contains 1 child
	if curr.Left == nil || curr.Right == nil {
		temp := curr.Left
		if temp == nil {
			temp = curr.Right
		}

		if parent.Left == curr {
			parent.Left = temp
		} else {
			parent.Right = temp
		}
	} else if curr.Left != nil && curr.Right != nil { // node contains 2 children

		temp := curr.Left
		prev := curr

		for temp.Right != nil { // get rightmost max on LST
			prev = temp
			temp = temp.Right
		}

		left := temp.Left
		curr.Val = temp.Val
		if prev.Left == temp {
			prev.Left = left
		} else {
			prev.Right = left
		}
	}

	if root.Val == k {

		if root.Left == nil {
			return root.Right
		}
		return root.Left
	}

	return root
}

func searchNode(root *TreeNode, k int) (*TreeNode, *TreeNode) {
	parent := new(TreeNode)

	temp := root

	for temp != nil {
		// may require equality check
		if temp.Val == k {
			return parent, temp
		} else if temp.Val < k {
			parent = temp
			temp = temp.Right
		} else {
			parent = temp
			temp = temp.Left
		}
	}

	return parent, nil
}
