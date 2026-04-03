/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		} else {
			right := dfs(root.Right)
			left :=  dfs(root.Left)
			res = max(res, right + left)
			return 1 + max(left, right)
		}
	}
	dfs(root)

	return res
}
