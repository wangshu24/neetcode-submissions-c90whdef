/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	} else {
		depth := max(maxDepth(root.Left), maxDepth(root.Right))
		return depth+1
	}

}
