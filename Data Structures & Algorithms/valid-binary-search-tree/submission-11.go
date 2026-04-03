/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	var res bool

	var dfs func(*TreeNode, int, int) bool
	dfs = func(temp *TreeNode, minN, maxN int) bool {
		if temp == nil {
			return true
		}
		if temp.Val <= minN || temp.Val >= maxN {
			return false
		}

		return dfs(temp.Left, minN, temp.Val) && dfs(temp.Right, temp.Val, maxN)
	}
	res = dfs(root, -1001, 1001)

	return res
}
