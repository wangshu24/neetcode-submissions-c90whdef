/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	balanced := true
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if left - right > 1 || right - left > 1 {
			balanced = false
		}
		return 1 + max(left,right)
	}
	dfs(root)
	return balanced
}

func abs(x int) int {
	if x < 0 {
		return -1
	}
	return x
}
