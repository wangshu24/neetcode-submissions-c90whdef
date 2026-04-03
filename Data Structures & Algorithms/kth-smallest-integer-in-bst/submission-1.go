/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    var res int
	count := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		} else if count == k {
			return
		}
		
		dfs(root.Left)
		count++
		if count == k  {
			res = root.Val
			return
		}
		dfs(root.Right)
		return
	}

	dfs(root)

	return res
}