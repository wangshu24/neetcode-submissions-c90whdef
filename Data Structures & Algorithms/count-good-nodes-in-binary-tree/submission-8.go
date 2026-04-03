/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    var good int

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, maxN int) {
		if root == nil {
			return
		}
		if root.Val >= maxN {
			good++
		}
		maxN = max(maxN, root.Val)
		dfs(root.Left, maxN)
		dfs(root.Right, maxN)
	}

	dfs(root, -101)

	return good
}
