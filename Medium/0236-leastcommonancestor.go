package Medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	path1 := make([]*TreeNode, 0)
	path2 := make([]*TreeNode, 0)

	preorder(root, p.Val, &path1)
	preorder(root, q.Val, &path2)
	//fmt.Println(path1, path2)
	for i := range path1 {
		for j := range path2 {

			if path1[i].Val == path2[j].Val {
				return path1[i]
			}
		}
	}

	return new(TreeNode)

}

func preorder(root *TreeNode, k int, result *[]*TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Val == k {
		*result = append(*result, root)
		return true
	}

	if preorder(root.Left, k, result) || preorder(root.Right, k, result) {
		*result = append(*result, root)
		return true
	}

	return false
}